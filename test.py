import subprocess

print(subprocess.run(["git", "add", "."]))
print(subprocess.run(["git", "commit", "-m", "\"Commitest\""]))