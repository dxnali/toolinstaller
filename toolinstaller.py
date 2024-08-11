#!/usr/bin/env python3

# ANSI colour
BOLD_BLACK = "\033[1;30m"
BOLD_RED = "\033[1;31m"
BOLD_GREEN = "\033[1;32m"
BOLD_YELLOW = "\033[1;33m"
BOLD_BLUE = "\033[1;34m"
BOLD_MAGENTA = "\033[1;35m"
BOLD_CYAN = "\033[1;36m"
BOLD_WHITE = "\033[1;37m"
RESET = "\033[0m"

def choose_superuser():
    while True:
        superuser = input(f"{BOLD_GREEN}[+] {BOLD_CYAN}Choose superuser (i.e. sudo, doas)\n===> {BOLD_BLUE}").strip().lower()
        if superuser in ["sudo", "doas"]:
            print(f"{BOLD_GREEN}[+] {BOLD_YELLOW}[{superuser}] {BOLD_CYAN}selected.{RESET}")
            return superuser
        else:
            print(f"{BOLD_RED}[-] {BOLD_BLUE}Invalid choice. Please choose 'sudo' or 'doas'.{RESET}")

def choose_pkgmnger():
    while True:
        pkgmnger = input(f"{BOLD_GREEN}[+] {BOLD_CYAN}Choose package manager (i.e. apt, dnf, pacman, apk)\n===> {BOLD_BLUE}").strip().lower()
        if pkgmnger in ["apt", "dnf", "pacman", "apk"]:
            print(f"{BOLD_GREEN}[+] {BOLD_YELLOW}[{pkgmnger}] {BOLD_CYAN}selected.{RESET}")
            return pkgmnger
        else:
            print(f"{BOLD_RED}[-] {BOLD_BLUE}Invalid choice. Please choose from 'apt', 'dnf', 'pacman', or 'apk'.{RESET}")

def add_packages():
    packages = []
    while not packages:
        package_string = input(f"{BOLD_GREEN}[+] {BOLD_CYAN}Add packages separated with commas (i.e. 'htop, firefox, alacritty')\n===> {BOLD_BLUE}").strip().lower()
        packages = [pkg.strip() for pkg in package_string.split(',') if pkg.strip()]
        if not packages:
            print(f"{BOLD_RED}[-] {BOLD_BLUE}No valid packages entered. Please try again.{RESET}")
    print(f"{BOLD_GREEN}[+] {BOLD_CYAN}Packages to install: {BOLD_YELLOW}{packages}{RESET}")
    return packages

def build_script(superuser, pkgmnger, addoption, packages):
    scriptname = input(f"{BOLD_GREEN}[+] {BOLD_CYAN}Enter bash script name without '.sh'\n===> {BOLD_BLUE}").strip().lower()
    scriptname = scriptname + '.sh'
    
    script = ['#!/bin/bash\n\n']

    for package in packages:
        command = f"{superuser} {pkgmnger} {addoption} {package}"
        if pkgmnger == "apt":
            command += " -y"
        script.append(command + "\n")

    try:
        with open(scriptname, 'w') as file:
            file.write(''.join(script))
        print(f"{BOLD_GREEN}[+] {BOLD_CYAN}Script written to {BOLD_YELLOW}[{scriptname}] {BOLD_CYAN}successfully.{RESET}")
    except Exception as e:
        print(f"{BOLD_RED}[-] {BOLD_BLUE}Error writing to file: {BOLD_YELLOW}{e}{RESET}")

def main():
    superuser = choose_superuser()
    pkgmnger = choose_pkgmnger()
    addoption = input(f"{BOLD_GREEN}[+] {BOLD_CYAN}Enter the method to add packages with your package manager (e.g., 'install' for apt, 'add' for apk)\n===> {BOLD_BLUE}").strip().lower()
    print(f"{BOLD_GREEN}[+] {BOLD_CYAN}Install option: {BOLD_YELLOW}[{addoption}]{RESET}")
    packagelist = add_packages()
    build_script(superuser, pkgmnger, addoption, packagelist)

if __name__ == "__main__":
    main()
