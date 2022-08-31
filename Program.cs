using System;

namespace Glow {
    class Program {
        public static void Main(string [] args){
            Console.WriteLine("Hello, World!");

            CommandExecutor.Command();  

            Console.WriteLine("End of the world"); 
            Console.ReadKey();
        }
    }
}

