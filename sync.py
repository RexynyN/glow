import subprocess

class Sync:
    def __init__(self, args):
        self.args = args
        pass

    def sync_down(self):

        subprocess.run(["git", "add", "."])
        subprocess.run(["git", "commit", "-m", self.args.m])
        subprocess.run(["git", "push", "-u", "origin", "main"])
        subprocess.run(["ls", "-l"])


if __name__ == "__main__":
    main()