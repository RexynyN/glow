def filerename(args):
    print(args)
    pass


# Logic if the script is called by itself
if __name__ == "__main__":
    import argparse
    from os import getcwd

    parser = argparse.ArgumentParser()
    # Adicionar os argumentos posicionais
    parser.add_argument("expression", type=str, help="String to replace the name of the files for")
    parser.add_argument("-d", "--dir", type=str, default=getcwd(), help="Directory of the folder to rename the files (current directory is default)")
    parser.add_argument("-s", "--start", type=int, default=0, help="Start of the naming sequence")
    args = parser.parse_args()

    filerename(args)
