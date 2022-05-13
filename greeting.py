def hi():
    print('안녕, 내일 봐')

def hello(msg):
    print(msg)

def bye():
    print('잘가')

print(f'__name__ 속성값은 {__name__}')

if __name__ == '__main__':
    print(hi())
    print(hello('오늘은 여기까지'))
    bye()
    
