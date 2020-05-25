print("How many kilometres is the distance?")
kms = float(input())
mls = kms/1.60934
mls = round(mls,3)
print(f"You said {kms}. In miles it is {mls}")