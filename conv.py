with open("123.txt") as f:
    with open("456.txt", "w") as f1:
        for line in f:
            line.encode('latin1').decode('cp1251').encode('UTF-8')
            f1.write(line) 
print("success")