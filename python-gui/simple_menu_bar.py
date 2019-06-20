import tkinter as tk
import sys

main_window = tk.Tk()
main_window.title("*** Simple Menu Bar ***")

value = tk.IntVar()
value.set(0)

label = tk.Label(main_window, textvariable=value)
label.pack()

menu_bar = tk.Menu(main_window)
main_window.configure(menu=menu_bar)

system_menu = tk.Menu(menu_bar)
menu_bar.add_cascade(menu=system_menu, label="System")

system_menu.add_radiobutton(label="One", variable=value, value=1)
system_menu.add_radiobutton(label="Two", variable=value, value=2)
system_menu.add_separator()
system_menu.add_command(label="Exit", command=sys.exit)

main_window.mainloop()