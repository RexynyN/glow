using Glow.Commands;
using System.Diagnostics;
using System.IO;
using System.Text.Json;
using System.Text.Json.Serialization;

namespace Glow
{
    class ImageData
    {
        private ImageDataOptions args;

        public ImageData(ImageDataOptions args)
        {
            this.args = args;
        }

        public void OrchestrateCommand()
        {
            if(args.Convert){

            }

            if(args.Revert){
                
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
}