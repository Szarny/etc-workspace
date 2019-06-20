import sys

import tkinter as tk
from tkinter import ttk


def change_label_value(buffer, value):
    return lambda: buffer.set(value)


def generate_style():
    ttk.Style().configure("btn.TButton", fg="#ffffff", bg="#000000")


def main():
    main_window = tk.Tk()
    main_window.title("Select Button")
    main_window.geometry("400x400")
    main_window.option_add("*Label.Background", "cyan")


    buffer = tk.StringVar()
    buffer.set("")
    label = ttk.Label(main_window, textvariable=buffer)
    label.pack()

    for v in range(5):
        btn = ttk.Button(main_window, text=str(v), style="btn.TButton", command=change_label_value(buffer, str(v)))
        btn.pack()

    main_window.mainloop()


if __name__ == '__main__':
    main()