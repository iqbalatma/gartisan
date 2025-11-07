package utils

import "fmt"

func PrintIntroduction() {
	// --- Bagian 1: Introduksi dan Deskripsi Singkat ---
	fmt.Println("---------------------------------------------------------")
	fmt.Println("✨ GARTISAN CLI - Command Line Tool For Go Developer ✨")
	fmt.Println("---------------------------------------------------------")
	fmt.Println("Gartisan is a CLI inspired by Laravel's Artisan tool")
	fmt.Println("designed to automate common development tasks in Golang.")
	fmt.Println()

	// --- Section 2: Basic Usage (Syntax) ---
	fmt.Println("## 🚀 Basic Usage (Syntax)")
	fmt.Println("Use the following format to run commands:")
	fmt.Println("gartisan <command> [arguments] [--flag]")
	fmt.Println()

	// --- Section 3: Usage example ---
	fmt.Println("## 💡 Examples")
	fmt.Println("gartisan serve --port 3000     // Runs the web server on port 3000")
	fmt.Println("gartisan migrate --rollback    // Reverts the last database migration")
	fmt.Println("gartisan make:controller management/user_controller  // Generates a new User controller file")
	fmt.Println()

	// --- Section 4: List available command ---
	fmt.Println("## 📋 Available Commands")
	fmt.Println("Perintah Utama:")
	fmt.Println("Generator Command (make):")
	fmt.Println("make:controller | Create new controller.")
	fmt.Println()

	// --- Bagian 5: Bantuan Lebih Lanjut ---
	fmt.Println("## ❓ Any question ? send email to iqbalatma@gmail.com")
	fmt.Println("\nEnjoy!")
}
