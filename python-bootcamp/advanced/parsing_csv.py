from csv import reader, DictReader


def parsing(file_name):
    with open(file_name, 'r') as file:
        data = file.read()
        parsed_lines = data.split('\n')
        parsed_lines.pop()
        parsed = [item.split(',') for item in parsed_lines]
        return parsed


print(parsing('../../original.csv'))

def pw_reader(file_name):
    with open(file_name) as file:
        csv_reader = reader(file)
        for row in csv_reader:
            print(row)


pw_reader('../../original.csv')


def pw_reader(file_name):
    with open(file_name) as file:
        csv_reader = DictReader(file)
        for row in csv_reader:
            print(row)


pw_reader('../../original.csv')
