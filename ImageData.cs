using Glow.Commands;
using System.Diagnostics;
using System.IO;

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
                            
        }

        private void CreateFile(List<string> values)
        {
            var json = JsonSerializer.Serialize(primes);
            File.WriteAllText(Path.Combine(Directory.GetCurrentDirectory(), "primes.json"), json);
        } 

    }
}