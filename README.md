# go_rng 1.1
by Thiago Jung  
https://github.com/Thiagojm/go_rng  
thiagojm1984@hotmail.com   
Written in Golang 1.19.4
-----------------------

# ABSTRACT

This application uses a type of TRNG - True Random Number Generator (TrueRNG) for data collection and statistical analysis for several purposes, including mind-matter interaction research. If you don't have the device, you can use the application to generate pseudo-random numbers aswell.  
 
It uses random numbers to collect and count the number of times the '1' bit appears in a series of user-defined size and interval. It will save the data in two files in the same directory as the application, one with a .csv extension and the other with a .bin extension (raw data to serve as control).

Afterwards, the data can be analyzed and compared with the number expected by chance (50%) and create a chart with a cumulative Z-Score. For that please use the RngKitPSG application at https://github.com/Thiagojm/RngKitPSG.  

# Supported OS:

1. Windows 10;  
2. Linux (Ubuntu / Debian-based/ Raspberry Pi OS).

# Supported Hardware:

1. TrueRNG and TrueRNGPro (https://ubld.it/);  
2. In built software pseudo-random number generator.  

# Installation

Windows INSTRUCTIONS
--------------------

1. Install Golang at https://go.dev/;

2. Open a command prompt (run the cmd command in Windows) and run:

`git clone github.com/Thiagojm/go_rng`

3. Change to the directory you created above and run: 

`go mod tidy` only once

4. Plug in a single TrueRNG V1, V2, V3, Pro, or ProV2

5. Run the application:

`go run .`

Linux Instructions
------------------------------------------

1. Install Golang at https://go.dev/;

2. Open a shell and run:

`git clone github.com/Thiagojm/go_rng`

3. Change to the directory you created above and run: 

`go mod tidy` only once

4. Plug in a single TrueRNG V1, V2, V3, Pro, or ProV2

5. Run the application:

`go run .`

Windows Drivers - if needed
--------------------

Choose from the Windows_Drivers folder the right driver for your hardware, right-click the TrueRNG.inf or TrueRNGpro.inf file and select Install. Follow the instructions for installation.

Linux Drivers - if needed
------------------------------------------

If needed, install the linux drivers as described in the README.md in the udev_rules folder.


# File Naming:

The file name contains important information about the collected data. The first part is the date and time of the collection, then the device used (trng for TrueRNG and pseudo for PseudoRNG), the number of bits per sample, the time between each sample in seconds. For example "20201011-142208_trng_s2048_i1": Collected on October 11, 2020 (20201011), at 14:22:08 (142208), TrueRNG device (trng), sample of 2048 bits (s2048) every 1 second (i1).

# Usage:

1. The application will ask you to enter the number of bits per sample. The default is 2048 bits. You can enter any divisible by 8. It will ask you to enter the time between each sample in seconds. The default is 1 second. You can enter any integer number.
2. Then the application will look for a TrueRNG device. If it finds one, it will start collecting data. If it doesn't find one, it will start collecting pseudo-random numbers (software based).
3. If you want to stop the collection, press Ctrl+C.
4. If you want, you can create a default.txt file in the same directory as the application and put the number of bits per sample and the time between each sample in seconds. The application will read the file and use the values in it. The file must have the following format; first line: number of bits per sample, second line: time between each sample in seconds. No spaces or other characters. For example (2048 bits per sample and 1 second between each sample):

`2048`  
`1`  

# Auto Start on Boot - Raspberry Pi OS

1- Install x-term:

`sudo apt-get install xterm`

2- If your username is pi, else change it accordingly:

`mkdir /home/pi/.config/autostart`

`nano /home/pi/.config/autostart/rng.desktop`

3- Paste the following lines (change the path to go_rng if needed)):

[Desktop Entry]  
Type=Application  
Name=RNG  
Exec=xterm -hold -e 'cd /home/pi/Desktop/go_rng/ && go run .'  

4- Now it shoud start the aplication everytime the system boot and start collecting.

