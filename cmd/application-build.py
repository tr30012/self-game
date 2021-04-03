import os

os.chdir('../src')
os.system('go build .')
os.system('move "web-game.exe" ../cmd')