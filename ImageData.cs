using Glow.Commands;
using System.Diagnostics;
using System.IO;
using System.Text.Json;
using System.Text.Json.Serialization;
using SixLabors.ImageSharp;

namespace Glow
{
    class ImageData
    {
        private ImageDataOptions args;

        public ImageData(ImageDataOptions args)
        {
            this.args = args;
        }

        public void Start()
        {
            if(args.Path == ".")
                args.Path = Directory.GetCurrentDirectory();

            switch(args.Action.ToLower())
            {
                case "convert":
                    ConvertImageData();
                    break;

                case "revert":
                    RevertImageData();
                    break;

                default: 
                    Console.WriteLine("No such action, try again.");
                    break;
            }
        }

        private void RevertImageData()
        {
            System.Console.WriteLine("Clocl");
        }

        private void ConvertImageData()
        {
            string [] whitelist = new string [] { "jpg", "jpeg", "jfif", "png", "webp", "gif" };
            List<FileInfo> pics = new List<FileInfo>();
            foreach(FileInfo file in new DirectoryInfo(args.Path).GetFiles())
            {
                if(whitelist.Any(extension => file.Name.EndsWith(extension)))
                    pics.Add(file);
            }

            
            int fileLength = 0;
            int fileIndex = 1; 
            foreach (var item in pics)
            {
                if(fileLength >= args.Max)
                {

                    fileIndex++;

                }

                string dataUri = Base64Image.Parse(item.Name, Image.Load(item.FullName)).ToDataUrl();
                fileLength++;

            }

        }

        public void ConvertImages()
        {
            if (!args.Dir)
                return;
                            
        }

        private void CreateFile(List<string> values)
        {
            var json = JsonSerializer.Serialize(values);
            File.WriteAllText(Path.Combine(Directory.GetCurrentDirectory(), "primes.json"), json);
        } 

    }

    class JsonImageDataSchema
    {
        public string Filename { get; set; }
        public string Content { get; set; }
        public string ContentType { get; set; }
    }
}