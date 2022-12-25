using System.Collections.Generic;

namespace Lab3
{
    public struct Directory
    {
        public string Name { get; }
        public bool IsDirectory { get; }
        public List<Directory> subdirectories;

        public Directory(string name_, bool isDirectory_)
        {
            Name = name_;
            IsDirectory = isDirectory_;
            subdirectories = new List<Directory>();
        }

        public void AddSubdirectory(Directory subdirectory_)
        {
            subdirectories.Add(subdirectory_);
        }
    }
}