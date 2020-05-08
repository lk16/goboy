package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"runtime/pprof"
	"syscall"
	"time"

	"github.com/faiface/mainthread"
	"github.com/sqweek/dialog"

	"fmt"

	"github.com/Humpheh/goboy/pkg/gb"
	"github.com/Humpheh/goboy/pkg/gb/io"
	"github.com/faiface/pixel/pixelgl"
)

// The version of GoBoy
var version = "develop"

const logo = `
    ______      ____
   / ____/___  / __ )____  __  __
  / / __/ __ \/ __  / __ \/ / / /
 / /_/ / /_/ / /_/ / /_/ / /_/ /
 \____/\____/_____/\____/\__, /
%23s /____/
`

var (
	mute    = flag.Bool("mute", false, "mute sound output")
	dmgMode = flag.Bool("dmg", false, "set to force dmg mode")

	cpuprofile   = flag.String("cpuprofile", "", "write cpu profile to file (debugging)")
	vsyncOff     = flag.Bool("disableVsync", false, "set to disable vsync (debugging)")
	stepThrough  = flag.Bool("stepthrough", false, "step through opcodes (debugging)")
	unlocked     = flag.Bool("unlocked", false, "if to unlock the cpu speed (debugging)")
	logfps       = flag.Bool("logfps", false, "log fps stats every second (debugging)")
	frontendName = flag.String("frontend", "pixelgl", "select frontend")
	noGraphics   = flag.Bool("nographics", false, "do not compute graphics")
)

func main() {
	flag.Parse()
	pixelgl.Run(start)
}

func start() {

	// Load the rom from the flag argument, or prompt with file select
	rom := getROM()

	// If the CPU profile flag is set, then setup the profiling
	if *cpuprofile != "" {
		startCPUProfiling()
		defer pprof.StopCPUProfile()
	}

	if *unlocked {
		*mute = true
	}

	// Print the logo and the run settings to the console
	fmt.Println(fmt.Sprintf(logo, version))
	fmt.Printf("APU: %v\nCGB: %v\nROM: %v\n", !*mute, !*dmgMode, rom)

	var opts []gb.GameboyOption
	if !*dmgMode {
		opts = append(opts, gb.WithCGBEnabled())
	}
	if !*mute {
		opts = append(opts, gb.WithSound())
	}
	if *noGraphics {
		opts = append(opts, gb.WithoutGraphics())
	}

	// Initialise the GameBoy with the flag options
	gameboy, err := gb.NewGameboy(rom, opts...)
	if err != nil {
		log.Fatal(err)
	}
	if *stepThrough {
		gameboy.Debug.OutputOpcodes = true
	}

	// Create the monitor for pixels
	enableVSync := !(*vsyncOff || *unlocked)

	frontendCreators := map[string]func() gb.IOBinding{
		"pixelgl": func() gb.IOBinding { return io.NewPixelsIOBinding(enableVSync) },
		"dummy":   func() gb.IOBinding { return &io.Dummy{} },
	}

	frontendCreator, ok := frontendCreators[*frontendName]

	if !ok {
		log.Fatalf("Could not find frontend named %s", *frontendName)
	}

	frontend := frontendCreator()
	startGBLoop(gameboy, frontend)
}

func startGBLoop(gameboy *gb.Gameboy, monitor gb.IOBinding) {
	frameTime := time.Second / gb.FramesSecond
	if *unlocked {
		frameTime = 1
	}

	var (
		frameTicker      = time.NewTicker(frameTime)
		lastStats        = time.Now()
		lastStatsFrames  int
		buttonProcessing time.Duration
		gameboyUpdating  time.Duration
		rendering        time.Duration
		frames           int
		cartName         string
	)

	if gameboy.IsGameLoaded() {
		cartName = gameboy.Memory.Cart.GetName()
	}

	// frames between logging
	const logInterval = time.Second

	for monitor.IsRunning() {

		<-frameTicker.C
		frames++

		start := time.Now()
		buttons := monitor.ButtonInput()
		gameboy.ProcessInput(buttons)
		buttonProcessing += time.Since(start)

		start = time.Now()
		_ = gameboy.Update()
		gameboyUpdating += time.Since(start)

		start = time.Now()
		monitor.Render(&gameboy.PreparedData)
		rendering += time.Since(start)

		if time.Since(lastStats) > logInterval {
			fps := float64(frames-lastStatsFrames) / logInterval.Seconds()
			title := fmt.Sprintf("GoBoy - %s (%d fps)", cartName, int(fps))
			monitor.SetTitle(title)

			if *logfps {
				log.Printf("| %4d fps | buttonProcessing = %6d µs/frame | gameboyUpdating = %6d µs/frame | rendering = %6d µs/frame",
					int(fps),
					int((1000000*buttonProcessing.Seconds())/fps),
					int((1000000*gameboyUpdating.Seconds())/fps),
					int((1000000*rendering.Seconds())/fps),
				)
			}

			frames = lastStatsFrames
			buttonProcessing = 0
			gameboyUpdating = 0
			rendering = 0
			lastStats = time.Now()
		}
	}
}

// Determine the ROM location. If the string in the flag value is empty then it
// should prompt the user to select a rom file using the OS dialog.
func getROM() string {
	rom := flag.Arg(0)
	if rom == "" {
		mainthread.Call(func() {
			var err error
			rom, err = dialog.File().
				Filter("GameBoy ROM", "zip", "gb", "gbc", "bin").
				Title("Load GameBoy ROM File").Load()
			if err != nil {
				os.Exit(1)
			}
		})
	}
	return rom
}

// Start the CPU profile to a the file passed in from the flag.
func startCPUProfiling() {
	log.Print("Starting CPU profile...")
	f, err := os.Create(*cpuprofile)
	if err != nil {
		log.Fatalf("Failed to create CPU profile: %v", err)
	}
	err = pprof.StartCPUProfile(f)
	if err != nil {
		log.Fatalf("Failed to start CPU profile: %v", err)
	}

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT)

	go func() {
		<-signalChan
		pprof.StopCPUProfile()
		os.Exit(0)
	}()
}
