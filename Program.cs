using System;
using System.IO;
using System.Drawing;
using static System.Net.Mime.MediaTypeNames;

namespace Glow {
    class Program {
        public static void Main(string [] args){

            var base64Img = new Base64Image 
            {
                FileContents = File.ReadAllBytes(@"C:\Users\017585631\Desktop\codes\glow\ibmhangout.jpg"),
                ContentType = "image/png"
            };

            string base64EncodedImg = base64Img.ToString();
            Console.WriteLine(base64EncodedImg);

            Console.WriteLine("Hello, World!");

            CommandExecutor.Command();
            ImageLib il = new ImageLib();

            Console.WriteLine("End of the world"); 
            Console.ReadKey();
        }
    }
}

