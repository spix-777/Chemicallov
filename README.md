# Chemicallov
 Chemical for Norway government list

 This Go program is a command-line tool that searches a list of banned chemicals for a specific chemical specified by the user. The list of banned chemicals is stored in a file called "tables.txt", which is created by scraping the HTML of a website that lists the banned chemicals. The user can also update the list of banned chemicals by using the "-u" flag when running the program.

The program begins by parsing command-line flags using the "flag" package. The "-u" flag is used to update the list of chemicals, and the "-w" flag is used to specify the chemical to search for. If the "-u" flag is set, the program calls the "updateChemicalList" function, which downloads the website containing the list of banned chemicals and parses the HTML to extract the tables containing the chemicals. The tables are then written to the "tables.txt" file.

If the "-w" flag is set and the "tables.txt" file exists, the program calls the "searchTable" function, which reads the contents of the file and searches for the specified chemical. If either the upper or lower case version of the chemical is found in the file, the program prints a message indicating that the chemical is banned. If the chemical is not found in the file, the program prints a message indicating that it is not banned.

Usage of lov:
  -u	Update list of chemicals
  -w string
    	Chemical for Norway government list (default "nil")
