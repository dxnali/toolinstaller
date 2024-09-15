package main

import (
        "bufio"
        "fmt"
        "os"
        "strings"
)

// ANSI color constants
const (
        BoldBlack   = "\033[1;30m"
        BoldRed     = "\033[1;31m"
        BoldGreen   = "\033[1;32m"
        BoldYellow  = "\033[1;33m"
        BoldBlue    = "\033[1;34m"
        BoldMagenta = "\033[1;35m"
        BoldCyan    = "\033[1;36m"
        BoldWhite   = "\033[1;37m"
        Reset       = "\033[0m"
)

func main() {
        superuser := chooseSuperuser()
        pkgManager := choosePkgManager()
        addOption := prompt(fmt.Sprintf("%s[+] %sEnter the method to add packages with your package manager (e.g., 'install' for apt, 'add' for apk)\n===> %s", BoldGreen, BoldCyan, BoldBlue))
        fmt.Printf("%s[+] %sInstall option: %s[%s]%s\n", BoldGreen, BoldCyan, BoldYellow, addOption, Reset)
        packageList := addPackages()
        buildScript(superuser, pkgManager, addOption, packageList)
}

func chooseSuperuser() string {
        for {
                superuser := prompt(fmt.Sprintf("%s[+] %sChoose superuser (i.e. sudo, doas)\n===> %s", BoldGreen, BoldCyan, BoldBlue))
                superuser = strings.ToLower(strings.TrimSpace(superuser))
                if superuser == "sudo" || superuser == "doas" {
                        fmt.Printf("%s[+] %s[%s] %sselected.%s\n", BoldGreen, BoldYellow, superuser, BoldCyan, Reset)
                        return superuser
                }
                fmt.Printf("%s[-] %sInvalid choice. Please choose 'sudo' or 'doas'.%s\n", BoldRed, BoldBlue, Reset)
        }
}

func choosePkgManager() string {
        for {
                pkgManager := prompt(fmt.Sprintf("%s[+] %sChoose package manager (i.e. apt, dnf, pacman, apk)\n===> %s", BoldGreen, BoldCyan, BoldBlue))
                pkgManager = strings.ToLower(strings.TrimSpace(pkgManager))
                if pkgManager == "apt" || pkgManager == "dnf" || pkgManager == "pacman" || pkgManager == "apk" {
                        fmt.Printf("%s[+] %s[%s] %sselected.%s\n", BoldGreen, BoldYellow, pkgManager, BoldCyan, Reset)
                        return pkgManager
                }
                fmt.Printf("%s[-] %sInvalid choice. Please choose from 'apt', 'dnf', 'pacman', or 'apk'.%s\n", BoldRed, BoldBlue, Reset)
        }
}

func addPackages() []string {
        var packages []string
        for {
                packageString := prompt(fmt.Sprintf("%s[+] %sAdd packages separated with commas (i.e. 'htop, firefox, alacritty')\n===> %s", BoldGreen, BoldCyan, BoldBlue))
                packages = parsePackages(packageString)

                if len(packages) == 0 {
                        fmt.Printf("%s[-] %sNo valid packages entered. Please try again.%s\n", BoldRed, BoldBlue, Reset)
                        continue
                }

                // Confirm package selection
                fmt.Printf("%s[+] %sPackages to install: %s%v%s\n", BoldGreen, BoldCyan, BoldYellow, packages, Reset)
                confirm := prompt(fmt.Sprintf("%s[?] %sDo you confirm these packages? (yes/no)\n===> %s", BoldYellow, BoldCyan, BoldBlue))
                if strings.ToLower(strings.TrimSpace(confirm)) == "yes" {
                        break
                } else {
                        fmt.Printf("%s[!] %sRe-enter the packages.\n", BoldRed, BoldCyan)
                }
        }
        return packages
}

func buildScript(superuser, pkgManager, addOption string, packages []string) {
        scriptName := prompt(fmt.Sprintf("%s[+] %sEnter bash script name:\n===> %s", BoldGreen, BoldCyan, BoldBlue))
        scriptName = strings.ToLower(strings.TrimSpace(scriptName))

        // Append ".sh" if not already present
        if !strings.HasSuffix(scriptName, ".sh") {
                scriptName += ".sh"
        }

        script := "#!/bin/bash\n\n"

        for _, pkg := range packages {
                command := fmt.Sprintf("%s %s %s %s", superuser, pkgManager, addOption, pkg)
                if pkgManager == "apt" {
                        command += " -y"
                }
                script += command + "\n"
        }

        err := os.WriteFile(scriptName, []byte(script), 0644)
        if err != nil {
                fmt.Printf("%s[-] %sError writing to file: %s%s%s\n", BoldRed, BoldBlue, BoldYellow, err.Error(), Reset)
        } else {
                fmt.Printf("%s[+] %sScript written to %s[%s] %ssuccessfully.%s\n", BoldGreen, BoldCyan, BoldYellow, scriptName, BoldCyan, Reset)
        }
}

func prompt(message string) string {
        fmt.Print(message)
        scanner := bufio.NewScanner(os.Stdin)
        scanner.Scan()
        return scanner.Text()
}

func parsePackages(packageString string) []string {
        var packages []string
        for _, pkg := range strings.Split(packageString, ",") {
                pkg = strings.TrimSpace(pkg)
                if pkg != "" {
                        packages = append(packages, pkg)
                }
        }
        return packages
}
