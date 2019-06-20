import tkinter as tk

def start():
    pass

root = tk.Tk()

level = tk.IntVar()
level.set(1)
action = tk.IntVar()
action.set(1)

menu_bar = tk.Menu(root)
root.configure(menu=menu_bar)

games = tk.Menu(menu_bar)
levels = tk.Menu(menu_bar)
menu_bar.add_cascade(menu=games, label="Games")
menu_bar.add_cascade(menu=levels, label="Levels")

games.add_command(label="Start", command=start)
games.add_separator()
games.add_radiobutton(label="First", variable=action, value=1)
games.add_radiobutton(label="Second", variable=action, value=2)
games.add_separator()
games.add_command(label="Exit", command=lambda: sys.exit())

levels.add_radiobutton(label = 'Level 1', variable = level, value = 1)
levels.add_radiobutton(label = 'Level 2', variable = level, value = 2)
levels.add_radiobutton(label = 'Level 3', variable = level, value = 3)

root.mainloop()