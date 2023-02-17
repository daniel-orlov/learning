from csv import DictReader, DictWriter


def update_users(f, l, nf, nl):
    with open("users.csv", "r+") as file:
        data = DictReader(file)

from csv import reader, writer

    def update_users(old_first, old_last, new_first, new_last):
        with open("users.csv") as csvfile:
            csv_reader = reader(csvfile)
            rows = list(csv_reader)

        count = 0
        with open("users.csv", "w") as csvfile:
            csv_writer = writer(csvfile)
            for row in rows:
                if row[0] == old_first and row[1] == old_last:
                    csv_writer.writerow([new_first, new_last])
                    count += 1
                else:
                    csv_writer.writerow(row)

        return "Users updated: {}.".format(count)
