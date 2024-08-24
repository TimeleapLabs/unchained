package ai

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/TimeleapLabs/unchained/internal/utils"

	"github.com/TimeleapLabs/unchained/internal/config"
	"github.com/TimeleapLabs/unchained/internal/precompiles/ai"
)

func startProcess(ctx context.Context, wg *sync.WaitGroup, env []string, cmdPath string, cmdArgs []string, cmdCwd string) {
	defer wg.Done()

	for {
		cmd := exec.CommandContext(ctx, cmdPath, cmdArgs...)
		cmd.Dir = cmdCwd
		cmd.Env = append(os.Environ(), env...)

		// Capture stdout and stderr
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		// Start the process
		if err := cmd.Start(); err != nil {
			log.Printf("Failed to start process: %v", err)
			return
		}

		log.Printf("Started process %d", cmd.Process.Pid)

		// Wait for the process to finish or be killed
		err := cmd.Wait()
		if ctx.Err() != nil {
			log.Printf("Process %d was killed", cmd.Process.Pid)
			return
		}

		if err != nil {
			log.Printf("Process %d exited with error: %v. Restarting...", cmd.Process.Pid, err)
		} else {
			log.Printf("Process %d exited successfully. Restarting...", cmd.Process.Pid)
		}

		// Sleep before restarting to avoid rapid restart loops
		time.Sleep(2 * time.Second)
	}
}

func StartServer(ctx context.Context) (*sync.WaitGroup, context.CancelFunc) {
	targetDir := filepath.Join(config.App.System.Home, "ai")

	// Extract Python files
	utils.Logger.Info("Extracting the plugin Python files...")
	if err := ai.ExtractPythonFiles(targetDir); err != nil {
		log.Fatalf("Failed to extract Python files: %v", err)
	}

	// Check if pyenv is installed
	_, err := exec.LookPath("pyenv")
	if err != nil {
		log.Fatal("pyenv not found in PATH")
	}

	// Install Python 3.8.10 using pyenv if not already installed
	checkPythonCmd := exec.Command("pyenv", "versions", "--bare")
	output, err := checkPythonCmd.Output()
	if err != nil || !strings.Contains(string(output), "3.8.10") {
		utils.Logger.Info("Installing Python 3.8.10...")
		installPythonCmd := exec.Command("pyenv", "install", "-s", "3.8.10")
		// installPythonCmd.Stdout = os.Stdout
		// installPythonCmd.Stderr = os.Stderr
		if err := installPythonCmd.Run(); err != nil {
			log.Fatalf("Failed to install Python 3.8.10: %v", err)
		}
	}

	// Select Python 3.8.10 as the local version
	utils.Logger.Info("Selecting Python 3.8.10...")
	selectPythonCmd := exec.Command("pyenv", "local", "3.8.10")
	selectPythonCmd.Dir = targetDir
	// selectPythonCmd.Stdout = os.Stdout
	// selectPythonCmd.Stderr = os.Stderr
	if err := selectPythonCmd.Run(); err != nil {
		log.Fatalf("Failed to select Python 3.8.10: %v", err)
	}

	// Get the path to the Python 3.8.10 interpreter
	pythonPathCmd := exec.Command("pyenv", "which", "python3.8")
	pythonPathCmd.Dir = targetDir
	pythonPath, err := pythonPathCmd.Output()
	if err != nil {
		log.Fatalf("Failed to get Python 3.8.10 path: %v", err)
	}
	pythonPathStr := strings.TrimSpace(string(pythonPath))

	// Create a virtual environment with Python 3.8.10 if not already created
	venvPath := filepath.Join(targetDir, "venv")
	if _, err := os.Stat(filepath.Join(venvPath, "bin", "python")); os.IsNotExist(err) {
		utils.Logger.Info("Creating virtual environment...")
		createVenvCmd := exec.Command(pythonPathStr, "-m", "venv", venvPath)
		// createVenvCmd.Stdout = os.Stdout
		// createVenvCmd.Stderr = os.Stderr
		if err := createVenvCmd.Run(); err != nil {
			log.Fatalf("Failed to create virtual environment: %v", err)
		}
	}

	// Activate the virtual environment
	activateScript := filepath.Join(venvPath, "bin", "activate")

	// Install dependencies if not already installed
	pipPath := filepath.Join(venvPath, "bin", "pip")
	freezeCmd := exec.Command("bash", "-c", fmt.Sprintf("source %s && exec %s freeze", activateScript, pipPath)) //nolint: gosec // This is a trusted command
	output, err = freezeCmd.Output()
	if err != nil || string(output) != string(ai.GetRequirementsFile()) {
		utils.Logger.Info("Installing dependencies...")
		requirementPath := filepath.Join(targetDir, "python_files", "requirements.txt")
		installDepsCmd := exec.Command("bash", "-c", fmt.Sprintf("source %s && %s install -r %s", activateScript, pipPath, requirementPath)) //nolint: gosec // This is a trusted command
		// installDepsCmd.Stdout = os.Stdout
		// installDepsCmd.Stderr = os.Stderr
		if err := installDepsCmd.Run(); err != nil {
			log.Fatalf("Failed to install dependencies: %v", err)
		}
	}

	// Set up the context and wait group for process management
	ctx, cancel := context.WithCancel(ctx)
	var wg sync.WaitGroup

	// Define environment variables
	env := []string{
		"HF_HOME=" + filepath.Join(targetDir, "hf_home"),
		"PYTHONWARNINGS=ignore",
	}

	// Start the process in a separate goroutine
	wg.Add(1)
	activateScript = filepath.Join("venv", "bin", "activate")
	mainPyPath := filepath.Join("python_files", "main.py")
	pythonCommand := fmt.Sprintf("source %s && exec python %s", activateScript, mainPyPath)
	go startProcess(ctx, &wg, env, "bash", []string{"-c", pythonCommand}, targetDir)

	return &wg, cancel
}
