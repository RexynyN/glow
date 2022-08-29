
import argparse

parser = argparse.ArgumentParser(description="A multiplatform toolbox for everything you need", prog="wings")


# Adicionar os argumentos posicionais
parser.add_argument("largura", type=int, help="Largura do terreno")
parser.add_argument("comprimento", type=int, help="Comprimento do terreno")
parser.add_argument("action")
args = parser.parse_args()


def calcula_area(largura, comprimento):
    return largura * comprimento

if __name__ == "__main__":
    print(f"A area do terreno é de {calcula_area(args.largura, args.comprimento)}m²")


# def main():
#     # Windows command to create a Scheduled Task for running Wings on every startup
#     "SCHTASKS /CREATE /SC ONSTART /TN \"OnStart Folder Sync\" /TR \"wings syncdown --all\""

# if __name__ == "__main__":
#     main()
    
    
