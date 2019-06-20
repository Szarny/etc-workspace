import sys

import tkinter
from tkinter import ttk

def exit():
    print("[*] Exit.")
    sys.exit(0)


def generate_config():
    ttk.Style().configure("label.TLabel", fg="#ff0000")
    ttk.Style().configure("button.TButton", fg="#ffffff", bg="000000")


def main():
    print("[*] Activate.")

    main_window = tkinter.Tk()
    main_window.title("Hello, Tkinter")
    main_window.geometry("600x600")

    generate_config()

    label = ttk.Label(main_window, text="Hello", style="label.TLabel")
    label.pack()

    button = ttk.Button(main_window, text='END', style="button.TButton", command=exit)
    button.pack()

    main_window.mainloop()


if __name__ == '__main__':
    main()