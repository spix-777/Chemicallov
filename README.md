# Chemical lov
 Chemical for Norway government list

 Chemical: a command-line tool for searching a list of banned chemicals in Norway!

 Are you worried about using a certain chemical in your lab or workplace? With Chemical, you can quickly find out if it's allowed by the Norwegian government.

 Just run lov.go with the "-w" flag followed by the chemical you want to check, and it will search the "tables.txt" file for the chemical's name. If it's found, Chemical will alert you that the chemical is not allowed. If it's not found, you're in the clear!

 But wait, what if the "tables.txt" file doesn't exist or is out of date? No problem! Just use the "-u" flag to update the list of chemicals by scraping the government's website. Chemical will handle all the hard work for you.

 So next time you're unsure about a chemical, just let Chemical do the searching for you. Happy experimenting!

Usage of Chemical:
  -u	Update list of chemicals
  -w string
    	Chemical for Norway government list (default "nil")
