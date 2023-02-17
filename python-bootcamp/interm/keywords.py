import pyfiglet

message = input('What message do you want to print? ')
clr = input('What color is the message? ')

pyfiglet.print_figlet(message, 'standard', clr.upper())
