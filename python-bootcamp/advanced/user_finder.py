from csv import DictReader


def find_user(first, last):
    with open("users.csv", "r") as f:
        data = DictReader(f)
        line_num = 0
        for line in data:
            line_num +=1
            if line["First Name"] == first:
                if line["Last Name"] == last:
                    return line_num
        return "Not Here not found."

print(find_user('Colt', 'Steele'))
