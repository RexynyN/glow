using SixLabors.ImageSharp;
using SixLabors.ImageSharp.Formats;


namespace Glow
{
    // Created by Jeremy Bell on StackOverflow (Modified for this project). 
    // Thread: https://stackoverflow.com/questions/21325661/convert-an-image-selected-by-path-to-base64-string
    public class Base64Image
    {
        public string Filename { get; set; }
        public string ContentType { get; set; }
        public byte[] FileContents { get; set; }
        public string DataUri { get; set; }

        public Base64Image(string filename, string contentType, byte[] fileContents, string dataUri)
        {
            Filename = filename;
            ContentType = contentType;
            FileContents = fileContents;
            DataUri = dataUri;
        }

        public static Base64Image Parse(string dataUri, string filename)
        {
            if (string.IsNullOrEmpty(dataUri))
                throw new ArgumentNullException(nameof(dataUri));

            int indexOfSemiColon = dataUri.IndexOf(";", StringComparison.OrdinalIgnoreCase);
            string dataLabel = dataUri.Substring(0, indexOfSemiColon);
            string contentType = dataLabel.Split(':').Last();
            var startIndex = dataUri.IndexOf("base64,", StringComparison.OrdinalIgnoreCase) + 7;
            var fileContents = dataUri.Substring(startIndex);
            var bytes = Convert.FromBase64String(fileContents);

            return new Base64Image(filename, contentType, bytes, dataUri);
        }

        public static Base64Image Parse(string filename, byte[] fileContent)
        {
            if (fileContent == null)
                throw new ArgumentNullException(nameof(fileContent));

            string dataUri = Convert.ToBase64String(fileContent);
            string contentType = "image/" + filename.Split(".")[1];
            
            return new Base64Image(filename, contentType, fileContent, dataUri);
        }

        public static Base64Image Parse(string filename, Image image)
        {
            if (image == null)
                throw new ArgumentNullException(nameof(image));

            string dataUri = image.ToBase64String(Image.DetectFormat(filename));
            byte[] fileContent = Convert.FromBase64String(dataUri);
            string contentType = "image/" + filename.Split(".")[1];
            
            return new Base64Image(filename, contentType, fileContent, dataUri);
        }

        public void ToImageFile()
        {
            using var image = Image.Load(FileContents);

            if(!Filename.EndsWith(ContentType.Split("/")[1]))
                Filename = Filename.Split(".")[0] + ContentType.Split("/")[1];

            image.Save(Filename);
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