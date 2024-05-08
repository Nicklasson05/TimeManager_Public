# TimeManager_Public
Time Manager is a program designed to make is easier to log time on issues in GitLab 

To use Time Manager you need to have a Access Token with the following scope
* api

----------------------------------------------INSTALLATION------------------------------------------------


Install for Linux:-------------------------------LINUX------------------------------------------------------

1. Extract content of 'TimeManager.tar.xz' into a temp folder 
2. Open a Terminal
3. $ cd YOUR-FOLDER
4. $ ./configure
5. $ make
6. $ make user-install

Install for Window:--------------------------WINDOWS------------------------------------------------

Extract content of 'TimeManager.zip' to get a working .exe file.
For system wide installation is it requierd to access Source Code.
Do the Following Guide to get access to it and then return here.
1. After following the guide you have access to fyne command tools
2. open the source code folder in a terminal and run
3. $ fyne install
4. And after that you have installed Time Manager on your system

----------------------------------------------SOURCE-CODE-------------------------------------------------

Access and run source code:
	Requierd för running:
		Golang, 
		C compiler 64/bit,
		Fyne 
  
Install Go:-------------------------------------GOLANG-------------------------------------------------------
	
	Note: Ubuntu users install go from website and not via snap 
	Ubuntu:
		1. Download the latest verion from go from this link
		   link: https://go.dev/dl/
		2. Extract the downloaded data into a folder
		
		Set GOPATH varuabels
		3. open terminal
		4. $ nano ~/.bashrc
		5. scroll to the bottom 
		6. enter following 
		7. export GOPATH=$HOME/go
		8. export PATH=$PATH:$GOPATH/bin 
		9. ctrl + x 
		10. Enter
		
		11. $ source ~/.bashrc 
		or restart terminal
		Sedan kolla så det e installerat med 
		12. $ go version 
		
	Windows:	
		1. Download Go for your operation system 
	   	   link: https://go.dev/dl/
	   	2. Then run installer
	
Install C compiler:---------------------------C-COMPILER----------------------------------------------------
	
	Ubuntu:
		Install gcc and g++ which is a common C compiler for Ubuntu 
		1. open a terminal and execute
		2. $ sudo apt update
		3. $ sudo apt install build-essential

	Windows:
		1. Install tdm64 10.3.0-2.exe which includes gcc and g++
		   link: https://jmeubank.github.io/tdm-gcc/download/
		2. Then run installer
	
Install Fyne:----------------------------------------FYNE-----------------------------------------------------
	
	Note: only install fyne after you have installed go and a C compiler
		1. Open termial and do following
		2. $ mkdir myapp
		3. $ cd myapp
		4. $ go mod init myapp
		
		Install latest version of fyne by
		5. $ go get fyne.io/fyne/v2@latest
		Install fyne helper tool by
		7. $ go install fyne.io/fyne/v2/cmd/fyne@latest
		
		If your on Ubuntu {
		  $ sudo apt-get install libgl1-mesa-dev xorg-dev
		}
	
		8. Log out of computer and sign in again 
	
	official installation guide by fyne:
		link: https://docs.fyne.io/started/
		
		
Check installation: ---------------------CHECK-INSTALLATION---------------------------------------------
	
	To see if everything is setup 
	Download the TimeManager_Public repo from github
		1. $ git clone https://github.com/Nicklasson05/TimeManager_Public.git
		2. In terminal navigate to source code folder and run 
		Check go:
		3. $ go version 
		Check C compiler 
		4. $ g++ --version
		Check fyne
		5. $ fyne version
		6. $ fyne
	
		Note: This guide was made 2023/05/02
----------------------------------------------------------------------------------------------------------	
	
	

