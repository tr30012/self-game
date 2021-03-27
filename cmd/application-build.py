import os

os.chdir('../src')
os.system('go build .')
os.system('move "self-game.exe" ../cmd')