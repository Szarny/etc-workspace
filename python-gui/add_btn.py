import sys

import tkinter as tk
from tkinter import ttk

buffer = None

def add():
    return lambda: buffer.set(int(buffer.get()) + 1)


def exit():
    return lambda: sys.exit(0)


def main():
    global buffer

    main_window = tk.Tk()
    main_window.title("Add Button")
    main_window.geometry("500x500")

    ttk.Style().configure("a.TLabel", fg="#000000")
    ttk.Style().configure("a.TButton", fg="#ffffff", bg="#000000")

    buffer = tk.StringVar()
    buffer.set("0")
    label = ttk.Label(main_window, style="a.TLabel", textvariable=buffer)
    label.pack()

    add_button = ttk.Button(main_window, style="a.TButton", text="ADD", command=add())
    add_button.pack()

    exit_button = ttk.Button(main_window, style="a.TButton", text="EXIT", command=lambda: sys.exit(0))
    exit_button.pack()

    main_window.mainloop()


if __name__ == '__main__':
    main()