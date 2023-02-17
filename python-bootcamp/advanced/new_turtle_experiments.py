from turtle import Turtle as t
from time import sleep as slp

bob = ''

def square(turtle, length):
    turtle = t()
    for i in range (4):
        turtle.fd(length)
        turtle.rt(90)
#square(bob, 500)
def polygon(turtle, length, n):
  turtle = t()
  for i in range (n):
      turtle.fd(length)
      turtle.rt(360/n)
#polygon(bob, 15, 37)
def circle(turtle, r):
  turtle = t()
  p = 3.14159265758939
  n = 360
  length = (2*p*r)/n
  for i in range (n):
      turtle.fd(length)
      turtle.rt(360/n)
#circle(bob, 200)
def polyline(turtle, n, length, ang):
    turtle = t()
    for i in range (n):
        turtle.fd(length)
        turtle.lt(ang)
#polyline(bob, 18, 150, 100)
# def flower(turtle, n, length):
#     turtle = t()
#     step = length/(n)
#     step_num = length
#     for i in range (n):
#         # for r in range(2):
#         for s in range (step_num):
#             turtle.fd(step)
#             turtle.lt(10)
#         turtle.lt(180-(360/n))
#     slp(3)
#     turtle.mainloop()
# flower(bob, 8, 100)
# def flower(turtle, n, length):
#     turtle = t()
#     # step = length/(n*n)
#     # step = length/(n*n*n)
#     step = 1
#     # step_num = length*2
#     step_num = 100
#     for i in range (int(n*1.5)):
#         for r in range(2):
#             for s in range (step_num):
#                 turtle.fd(step)
#                 turtle.lt(5)
#             # turtle.lt(180-(360/n)) 120
#             turtle.lt(120)
#     slp(30)

def flower(turtle, n, length):
    turtle = t()
    angle_1 = 50
    for i in range (n):
        turtle.fd(length)
        turtle.lt(angle_1)
        turtle.fd(length)
        turtle.lt(180 - angle_1)
        turtle.fd(length)
        turtle.lt(angle_1)
        turtle.fd(length)
        turtle.rt(angle_1)
    slp(10)
    turtle.mainloop()
flower(bob, 6, 25)

