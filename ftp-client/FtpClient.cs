using System;
using System.IO;
using System.Net;
using System.Text;

namespace Lab3
{
    class FtpClient
    {
        readonly static int maxDeep = 3;

        readonly string host;
        readonly string username;
        readonly string password;
        Directory rootDirectory;

        public FtpClient(string host_, string username_, string password_)
        {
            host = host_;
            username = username_;
            password = password_;
            rootDirectory = new Directory(host_, true);
        }

        public Directory GetDirectoryStructure()
        {
            Start(host, username, password, rootDirectory, 1);
            return rootDirectory;
        }

        private void Start(string host, string username_, string password_, Directory rootDirectory_, int deep_)
        {
            if (deep_ > maxDeep)
                return;

            FtpWebRequest ftpWebRequest = (FtpWebRequest)WebRequest.Create("ftp://" + host);

            /* ListDirectoryDetails: 
             * команда LIST, 
             * возвращает подробный список файлов на FTP-сервере: */
            ftpWebRequest.Method = WebRequestMethods.Ftp.ListDirectoryDetails;
            ftpWebRequest.Credentials = new NetworkCredential(username_, password_);

            // при переопределении во вложенном классе возвращает ответ на интернет-запрос:
            FtpWebResponse ftpWebResponse = (FtpWebResponse)ftpWebRequest.GetResponse();

            /* реализует TextReader, 
             * который считывает символы из потока байтов в определённой кодировке: */
            StreamReader streamReader = new StreamReader(ftpWebResponse.GetResponseStream());

            /* получает значение, 
             * определяющее находится ли позиция текущего потока в конце потока: */
            if (!streamReader.EndOfStream)
            {
                /* cчитывает все символы, 
                начиная с текущей позиции до конца потока
                
                получим:
                respo "
                -rwxr-x---    1 1227     1000           18 Nov 20  2015 .bash_logout\r\n
                -rwxr-x---    1 1227     1000          193 Nov 20  2015 .bash_profile\r\n
                -rwxr-x---    1 1227     1000          231 Nov 20  2015 .bashrc\r\n
                drwxr-x---    2 1227     1000         4096 Nov 21  2016 .log\r\n
                drwxr-x---    4 1227     1000         4096 Dec 07  2016 htdocs\r\n
                dr-xr-xr--    2 1227     1000         4096 Dec 21  2016 test.test\r\n
                drwxr-x---    2 1227     1000         4096 Nov 21  2016 tmp\r\n
                " string 
                */
                string stringResponse = streamReader.ReadToEnd();

                char[] separator = new char[] { '\r', '\n' };

                /* массив, 
                 * элементы которого содержат подстроки из этого экземпляра, 
                 * разделённые символами из separator: */
                string[] lines = stringResponse.Split(separator, StringSplitOptions.RemoveEmptyEntries);
                /* теперь имеем lines:
                0   respo "
                    -rwxr-x---    1 1227     1000           18 Nov 20  2015 .bash_logout
                1   -rwxr-x---    1 1227     1000          193 Nov 20  2015 .bash_profile
                2   -rwxr-x---    1 1227     1000          231 Nov 20  2015 .bashrc
                3   drwxr-x---    2 1227     1000         4096 Nov 21  2016 .log
                4   drwxr-x---    4 1227     1000         4096 Dec 07  2016 htdocs
                5   dr-xr-xr--    2 1227     1000         4096 Dec 21  2016 test.test
                6   drwxr-x---    2 1227     1000         4096 Nov 21  2016 tmp
                7   " string 
                 */

                foreach (string line in lines)
                {
                    if (line[line.Length - 1] != '.') // если не ссылки на каталоги
                    {
                        string directoryName = GetDirectoryName(line);

                        if (line[0] == '-') // если файл
                            rootDirectory_.AddSubdirectory(new Directory(directoryName, false));

                        if (line[0] == 'd') // если папка
                        {
                            Directory subdirectory = new Directory(directoryName, true);

                            // рекурсия:
                            Start(host + "/" + directoryName, username_, password, subdirectory, deep_++);

                            rootDirectory_.AddSubdirectory(subdirectory);
                        }
                    }
                }
            }
        }

        private string GetDirectoryName(string line)
        {
            /* пусть есть line:
             -rwxr-x---     1     1227     1000     193     Nov     20     2015     .bash_profile 
             */
            string[] parts = line.Split(new char[] { ' ' }, StringSplitOptions.RemoveEmptyEntries);
            /* теперь имеем parts:
             0      -rwxr-x---     
             1      1     
             2      1227     
             3      1000    
             4      193     
             5      Nov     
             6      20     
             7      2015     
             8  .bash_profile 
             */

            if (parts.Length <= 9)  // в таком случае последний элемент
                return parts[parts.Length - 1];
            else                    // если в имени есть пробелы              
            {
                StringBuilder name = new StringBuilder();

                name.Append(parts[8]);
                for (int i = 9; i < parts.Length; i++)
                    name.Append(" " + parts[i]);

                return name.ToString();
            }
        }
    }
}