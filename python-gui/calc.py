import sys
import os
import tkinter as tk

flag = "CTF{FAOFJIAJFEOJFIOAOJF}"

def calc(e):
    if buffer.get():
        buffer.set(str(eval(buffer.get())))

main_window = tk.Tk()
main_window.title("Calc")
main_window.geometry("400x400")

buffer = tk.StringVar()
buffer.set("")

entry = tk.Entry(main_window, textvariable=buffer)
entry.pack()
entry.bind("<Return>", calc)

main_window.mainloop()