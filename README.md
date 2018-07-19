## Auto changing wallpaper using time of day

High Sierra ships with utility that allows people to change wallpaper based
on time of day. It has a cool transition, which is rather advanced, but the core functionality of changing wallpaper can be implemented using a cron job. I plan to implement this using GO Lang.

	the first step would be to make a go program which on running will
	update the wallpaper entry.

	the second step will be to make this program, run using cron job.

	the third setp will be make this cronjob load on startup

	fourth step to make a script which will automate installation.


	It will be nice to introduce a config file which keeps three catagories

	1. day 
	2. night
	3. random

	day: wallpaper which is to be shown during day.

	night: wallpaper which is to be shown during night.

	random: folder, which will have wallpaper set to be 
	cycled.

	4. startOfDay: time when day starts.
	5. endOfDay: time when day ends.
	6. cycleDuration: cycle duration for wallpaper in random.
	7. autoload: yes/no if this flag is set then, this program will
	run automatically at next startup. If this is unset and applied.
	this program will not run unless loaded by user.
