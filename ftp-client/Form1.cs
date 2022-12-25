using System;
using System.Windows.Forms;
using System.Drawing;

namespace Lab3
{
    public partial class Form1 : Form
    {
        public Form1()
        {
            InitializeComponent();
        }

        private void buttonStart_Click(object sender, EventArgs e)
        {
            try
            {
                FtpClient ftpClient = new FtpClient(textBoxHost.Text, textBoxUsername.Text, textBoxPassword.Text);
                Directory DirectoryStructure = ftpClient.GetDirectoryStructure();

                /* TreeView.Nodes: 
                 * получает коллекцию узлов дерева, 
                 * которая назначена элементу управления иерархического представления: */
                treeView.Nodes.Add(DirectoryStructure.Name);
                BuildTreeView(DirectoryStructure, treeView.Nodes[0]);
            }
            catch (Exception exc)
            {
                MessageBox.Show(exc.Message);
            }
        }

        private void BuildTreeView(Directory directory_, TreeNode treeNode_)
        {
            foreach (var subdirectory in directory_.subdirectories)
            {
                treeNode_.Nodes.Add(subdirectory.Name);

                if (subdirectory.IsDirectory)
                    BuildTreeView(subdirectory, treeNode_.Nodes[treeNode_.Nodes.Count - 1]);
                else
                    treeNode_.Nodes[treeNode_.Nodes.Count - 1].ForeColor = Color.Red;
            }
        }
    }
}