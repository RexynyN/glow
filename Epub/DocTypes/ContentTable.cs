using System;
using System.Collections.Generic;
using System.Data.SqlTypes;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using System.Xml.Serialization;

namespace Glow.Epub.DocTypes
{
    class ContentTable : IEpubDocument
    {
        public void OutputFile(string path)
        {
            throw new NotImplementedException();
        }
    }

    // Parse the toc.ncx using this shit 
    //https://docs.microsoft.com/en-us/dotnet/standard/serialization/examples-of-xml-serialization?source=recommendations


    [XmlRoot("ncx", Namespace = "http://www.daisy.org/z3986/2005/ncx/", IsNullable = false)]
    class Ncx
    {
        [XmlArray("head")]
        public Meta[] meta;

        public class Meta
        {
            [XmlAttribute]
            public string name { get; set; }
            [XmlAttribute]
            public string content { get; set; }
        }


        [XmlArray("navMap")]
        public navPoint[] navPoint;

        public class navPoint
        {
            string content 
        }
    }

    
}
