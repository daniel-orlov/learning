#Examples

    # divide(4,2)  2
    # divide([],"1")  "Please provide two integers or floats"
    # divide(1,0)  "Please do not divide by zero"

def divide(num1,num2):
    try:
        return num1/num2
    except (TypeError or ValueError):
        print("Please provide two integers or floats")
    except ZeroDivisionError:
        print("Please do not divide by zero")
