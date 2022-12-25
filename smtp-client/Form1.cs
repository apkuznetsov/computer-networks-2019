using System;
using System.Windows.Forms;

namespace Lab4
{
    public partial class Form1 : Form
    {
        public Form1()
        {
            InitializeComponent();
        }

        private void attachButton_Click(object sender, EventArgs e)
        {
            openFileDialog1.ShowDialog();
        }

        private void openFileDialog1_FileOk(object sender, System.ComponentModel.CancelEventArgs e)
        {
            attachmentTextBox.Text = openFileDialog1.FileName;
        }

        private void sendButton_Click(object s, EventArgs e)
        {
            try
            {
                SendMessage.Send(
                    fromTextBox.Text,
                    toTextBox.Text,
                    subjectTextBox.Text,
                    bodyRichTextBox.Text,
                    attachmentTextBox.Text,
                    hostTextBox.Text,
                    int.Parse(portTextBox.Text),
                    passwordTextBox.Text);

                MessageBox.Show("сообщение отправлено");
            }
            catch (Exception exc)
            {
                MessageBox.Show(exc.Message);
            }
        }
    }
}