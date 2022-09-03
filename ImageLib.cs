using System;
using System.IO;
using SixLabors.ImageSharp;
using SixLabors.ImageSharp.Processing;

namespace Glow
{

    internal class ImageLib
    {
        public ImageLib()
        {
        }

        public string ToDataUrl(string path)
        {
            byte[] imageArray = File.ReadAllBytes(path);
            return Convert.ToBase64String(imageArray);
        }

        //public void ConvertDataUrl(string dataurl)
        //{
        //    var image = Image.FromStream(new MemoryStream(Convert.FromBase64String(dataurl)));
        //    image.Save();
        //}
    }
}