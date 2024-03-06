# TimeManager_Public
A time tracking application that is designed around GitLabs time tacking system with issues:

Programed with Go ( this repotitory holds the Source code ) 
Second version(V2.5) is out for testing purposes

PS
Don't be afraid of giving feedback and report bugs 

IMPORTANT INFORMATION:

In the options menu you enter your own API key from gitlab
don't forget to save

LOGNING INSTRUKTIONS:

Play button:
  Loggs when you arrive at work for exampel, and should always be pressed when starting the application on a new day 

Record button:
  Is used in tandem with the entry below it to select a task that is about to be logged
    It register time based on the current time and the latest archived time in the log

Pause Button:
  It is used when you take a break or go on lunch to track how mutch salcking you have done in your work day

Stop Button:
  when this is pressed you stop the status texts timer in the down left corner 
  and should be used when you end you'r shift to register how long you have worked and other stats 
  Is also lockes you from using the logg buttons

TO UPLOAD YOU TIMES TO GITLAB:

  its realy simpel to register your work time 
  1. first step is to check in the dates in the top right that you want to upload
  2. then navigate to the send manu in the top left 
  3. then you can se what time will be logged to whish issue in your project 
  4. then press the Upload button in the top left corner

EDIT THE LOG
  you can edit the log through the big entry and then press the save button in the top left of the entry
  NOTE: its very sensitiv to utside input so be careful and try to follow the syntax on the appliaction
