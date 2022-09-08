using System;
using System.Linq;
using SixLabors.ImageSharp;

namespace Glow
{
    // Created by Jeremy Bell on StackOverflow (Modified for this project). 
    // Thread: https://stackoverflow.com/questions/21325661/convert-an-image-selected-by-path-to-base64-string
    public class Base64Image
    {
        public string ContentType { get; set; }
        public byte[] FileContents { get; set; }

        public Base64Image(string contentType, byte[] fileContents)
        {
            ContentType = contentType;
            FileContents = fileContents;
        }

        public Base64Image(string dataUri)
        {
            
        }

        public static Base64Image Parse(string base64Content)
        {
            if (string.IsNullOrEmpty(base64Content))
            {
                throw new ArgumentNullException(nameof(base64Content));
            }

            int indexOfSemiColon = base64Content.IndexOf(";", StringComparison.OrdinalIgnoreCase);
            string dataLabel = base64Content.Substring(0, indexOfSemiColon);
            string contentType = dataLabel.Split(':').Last();
            var startIndex = base64Content.IndexOf("base64,", StringComparison.OrdinalIgnoreCase) + 7;
            var fileContents = base64Content.Substring(startIndex);
            var bytes = Convert.FromBase64String(fileContents);

            return new Base64Image(contentType, bytes);
        }

        public void ToImageFile(string content)
        {
            byte[] data = Convert.FromBase64String(content);
            using var image = Image.Load(data);
        }   

        public string ToDataUrl()
        {
            return ToString();
        }

        public override string ToString()
        {
            return $"data:{ContentType};base64,{Convert.ToBase64String(FileContents)}";
        }
    }
}