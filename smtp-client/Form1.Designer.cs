namespace Lab4
{
    partial class Form1
    {
        /// <summary>
        /// Требуется переменная конструктора.
        /// </summary>
        private System.ComponentModel.IContainer components = null;

        /// <summary>
        /// Освободить все используемые ресурсы.
        /// </summary>
        /// <param name="disposing">истинно, если управляемый ресурс должен быть удален; иначе ложно.</param>
        protected override void Dispose(bool disposing)
        {
            if (disposing && (components != null))
            {
                components.Dispose();
            }
            base.Dispose(disposing);
        }

        #region Код, автоматически созданный конструктором форм Windows

        /// <summary>
        /// Обязательный метод для поддержки конструктора - не изменяйте
        /// содержимое данного метода при помощи редактора кода.
        /// </summary>
        private void InitializeComponent()
        {
            this.recipientsGroupBox = new System.Windows.Forms.GroupBox();
            this.attachButton = new System.Windows.Forms.Button();
            this.sendButton = new System.Windows.Forms.Button();
            this.subjectLabel = new System.Windows.Forms.Label();
            this.attachmentTextBox = new System.Windows.Forms.TextBox();
            this.subjectTextBox = new System.Windows.Forms.TextBox();
            this.attachLabel = new System.Windows.Forms.Label();
            this.bodyRichTextBox = new System.Windows.Forms.RichTextBox();
            this.toLabel = new System.Windows.Forms.Label();
            this.toTextBox = new System.Windows.Forms.TextBox();
            this.fromLabel = new System.Windows.Forms.Label();
            this.fromTextBox = new System.Windows.Forms.TextBox();
            this.senderGroupBox = new System.Windows.Forms.GroupBox();
            this.portTextBox = new System.Windows.Forms.TextBox();
            this.portLabel = new System.Windows.Forms.Label();
            this.passwordLabel = new System.Windows.Forms.Label();
            this.passwordTextBox = new System.Windows.Forms.TextBox();
            this.hostTextBox = new System.Windows.Forms.TextBox();
            this.hostLabel = new System.Windows.Forms.Label();
            this.openFileDialog1 = new System.Windows.Forms.OpenFileDialog();
            this.recipientsGroupBox.SuspendLayout();
            this.senderGroupBox.SuspendLayout();
            this.SuspendLayout();
            // 
            // recipientsGroupBox
            // 
            this.recipientsGroupBox.Controls.Add(this.attachButton);
            this.recipientsGroupBox.Controls.Add(this.sendButton);
            this.recipientsGroupBox.Controls.Add(this.subjectLabel);
            this.recipientsGroupBox.Controls.Add(this.attachmentTextBox);
            this.recipientsGroupBox.Controls.Add(this.subjectTextBox);
            this.recipientsGroupBox.Controls.Add(this.attachLabel);
            this.recipientsGroupBox.Controls.Add(this.bodyRichTextBox);
            this.recipientsGroupBox.Controls.Add(this.toLabel);
            this.recipientsGroupBox.Controls.Add(this.toTextBox);
            this.recipientsGroupBox.Font = new System.Drawing.Font("Microsoft Sans Serif", 12F, System.Drawing.FontStyle.Regular, System.Drawing.GraphicsUnit.Point, ((byte)(204)));
            this.recipientsGroupBox.Location = new System.Drawing.Point(263, 11);
            this.recipientsGroupBox.Margin = new System.Windows.Forms.Padding(2);
            this.recipientsGroupBox.Name = "recipientsGroupBox";
            this.recipientsGroupBox.Padding = new System.Windows.Forms.Padding(2);
            this.recipientsGroupBox.Size = new System.Drawing.Size(400, 380);
            this.recipientsGroupBox.TabIndex = 0;
            this.recipientsGroupBox.TabStop = false;
            this.recipientsGroupBox.Text = "письмо";
            // 
            // attachButton
            // 
            this.attachButton.Font = new System.Drawing.Font("Microsoft Sans Serif", 11.25F, System.Drawing.FontStyle.Regular, System.Drawing.GraphicsUnit.Point, ((byte)(204)));
            this.attachButton.Location = new System.Drawing.Point(92, 99);
            this.attachButton.Name = "attachButton";
            this.attachButton.Size = new System.Drawing.Size(107, 23);
            this.attachButton.TabIndex = 19;
            this.attachButton.Text = "открыть";
            this.attachButton.UseVisualStyleBackColor = true;
            this.attachButton.Click += new System.EventHandler(this.attachButton_Click);
            // 
            // sendButton
            // 
            this.sendButton.Font = new System.Drawing.Font("Microsoft Sans Serif", 12F, System.Drawing.FontStyle.Underline, System.Drawing.GraphicsUnit.Point, ((byte)(204)));
            this.sendButton.Location = new System.Drawing.Point(4, 350);
            this.sendButton.Margin = new System.Windows.Forms.Padding(2);
            this.sendButton.Name = "sendButton";
            this.sendButton.Size = new System.Drawing.Size(392, 26);
            this.sendButton.TabIndex = 28;
            this.sendButton.Text = "отправить";
            this.sendButton.Click += new System.EventHandler(this.sendButton_Click);
            // 
            // subjectLabel
            // 
            this.subjectLabel.AutoSize = true;
            this.subjectLabel.Font = new System.Drawing.Font("Microsoft Sans Serif", 11.25F, System.Drawing.FontStyle.Regular, System.Drawing.GraphicsUnit.Point, ((byte)(204)));
            this.subjectLabel.Location = new System.Drawing.Point(5, 67);
            this.subjectLabel.Name = "subjectLabel";
            this.subjectLabel.Size = new System.Drawing.Size(46, 18);
            this.subjectLabel.TabIndex = 25;
            this.subjectLabel.Text = "тема:";
            // 
            // attachmentTextBox
            // 
            this.attachmentTextBox.Font = new System.Drawing.Font("Microsoft Sans Serif", 10.2F, System.Drawing.FontStyle.Regular, System.Drawing.GraphicsUnit.Point, ((byte)(204)));
            this.attachmentTextBox.Location = new System.Drawing.Point(205, 99);
            this.attachmentTextBox.Name = "attachmentTextBox";
            this.attachmentTextBox.Size = new System.Drawing.Size(190, 23);
            this.attachmentTextBox.TabIndex = 15;
            // 
            // subjectTextBox
            // 
            this.subjectTextBox.Font = new System.Drawing.Font("Microsoft Sans Serif", 10.2F, System.Drawing.FontStyle.Regular, System.Drawing.GraphicsUnit.Point, ((byte)(204)));
            this.subjectTextBox.ForeColor = System.Drawing.SystemColors.AppWorkspace;
            this.subjectTextBox.Location = new System.Drawing.Point(92, 70);
            this.subjectTextBox.Name = "subjectTextBox";
            this.subjectTextBox.Size = new System.Drawing.Size(303, 23);
            this.subjectTextBox.TabIndex = 24;
            this.subjectTextBox.Text = "без темы";
            // 
            // attachLabel
            // 
            this.attachLabel.AutoSize = true;
            this.attachLabel.Font = new System.Drawing.Font("Microsoft Sans Serif", 11.25F, System.Drawing.FontStyle.Regular, System.Drawing.GraphicsUnit.Point, ((byte)(204)));
            this.attachLabel.Location = new System.Drawing.Point(5, 96);
            this.attachLabel.Name = "attachLabel";
            this.attachLabel.Size = new System.Drawing.Size(81, 18);
            this.attachLabel.TabIndex = 14;
            this.attachLabel.Text = "вложение:";
            // 
            // bodyRichTextBox
            // 
            this.bodyRichTextBox.Font = new System.Drawing.Font("Times New Roman", 10.8F, System.Drawing.FontStyle.Regular, System.Drawing.GraphicsUnit.Point, ((byte)(204)));
            this.bodyRichTextBox.ForeColor = System.Drawing.SystemColors.AppWorkspace;
            this.bodyRichTextBox.Location = new System.Drawing.Point(5, 128);
            this.bodyRichTextBox.Name = "bodyRichTextBox";
            this.bodyRichTextBox.Size = new System.Drawing.Size(390, 217);
            this.bodyRichTextBox.TabIndex = 23;
            this.bodyRichTextBox.Text = "текст сообщения";
            // 
            // toLabel
            // 
            this.toLabel.AutoSize = true;
            this.toLabel.Font = new System.Drawing.Font("Microsoft Sans Serif", 11.25F, System.Drawing.FontStyle.Regular, System.Drawing.GraphicsUnit.Point, ((byte)(204)));
            this.toLabel.Location = new System.Drawing.Point(5, 21);
            this.toLabel.Name = "toLabel";
            this.toLabel.Size = new System.Drawing.Size(120, 18);
            this.toLabel.TabIndex = 22;
            this.toLabel.Text = "кому, эл. почта:";
            // 
            // toTextBox
            // 
            this.toTextBox.Font = new System.Drawing.Font("Microsoft Sans Serif", 11.25F, System.Drawing.FontStyle.Regular, System.Drawing.GraphicsUnit.Point, ((byte)(204)));
            this.toTextBox.Location = new System.Drawing.Point(4, 41);
            this.toTextBox.Margin = new System.Windows.Forms.Padding(2);
            this.toTextBox.Name = "toTextBox";
            this.toTextBox.Size = new System.Drawing.Size(392, 24);
            this.toTextBox.TabIndex = 1;
            // 
            // fromLabel
            // 
            this.fromLabel.AutoSize = true;
            this.fromLabel.Font = new System.Drawing.Font("Microsoft Sans Serif", 11.25F, System.Drawing.FontStyle.Regular, System.Drawing.GraphicsUnit.Point, ((byte)(204)));
            this.fromLabel.Location = new System.Drawing.Point(5, 113);
            this.fromLabel.Name = "fromLabel";
            this.fromLabel.Size = new System.Drawing.Size(137, 18);
            this.fromLabel.TabIndex = 9;
            this.fromLabel.Text = "от кого, эл. почта:";
            // 
            // fromTextBox
            // 
            this.fromTextBox.Font = new System.Drawing.Font("Microsoft Sans Serif", 11.25F, System.Drawing.FontStyle.Regular, System.Drawing.GraphicsUnit.Point, ((byte)(204)));
            this.fromTextBox.Location = new System.Drawing.Point(4, 133);
            this.fromTextBox.Margin = new System.Windows.Forms.Padding(2);
            this.fromTextBox.Name = "fromTextBox";
            this.fromTextBox.Size = new System.Drawing.Size(238, 24);
            this.fromTextBox.TabIndex = 0;
            this.fromTextBox.Text = "test@mail.ru";
            // 
            // senderGroupBox
            // 
            this.senderGroupBox.Controls.Add(this.portTextBox);
            this.senderGroupBox.Controls.Add(this.portLabel);
            this.senderGroupBox.Controls.Add(this.passwordLabel);
            this.senderGroupBox.Controls.Add(this.passwordTextBox);
            this.senderGroupBox.Controls.Add(this.hostTextBox);
            this.senderGroupBox.Controls.Add(this.hostLabel);
            this.senderGroupBox.Controls.Add(this.fromLabel);
            this.senderGroupBox.Controls.Add(this.fromTextBox);
            this.senderGroupBox.Font = new System.Drawing.Font("Microsoft Sans Serif", 12F, System.Drawing.FontStyle.Regular, System.Drawing.GraphicsUnit.Point, ((byte)(204)));
            this.senderGroupBox.Location = new System.Drawing.Point(11, 11);
            this.senderGroupBox.Margin = new System.Windows.Forms.Padding(2);
            this.senderGroupBox.Name = "senderGroupBox";
            this.senderGroupBox.Padding = new System.Windows.Forms.Padding(2);
            this.senderGroupBox.Size = new System.Drawing.Size(248, 380);
            this.senderGroupBox.TabIndex = 1;
            this.senderGroupBox.TabStop = false;
            this.senderGroupBox.Text = "отправитель";
            // 
            // portTextBox
            // 
            this.portTextBox.Font = new System.Drawing.Font("Microsoft Sans Serif", 10.2F, System.Drawing.FontStyle.Regular, System.Drawing.GraphicsUnit.Point, ((byte)(204)));
            this.portTextBox.Location = new System.Drawing.Point(5, 87);
            this.portTextBox.Name = "portTextBox";
            this.portTextBox.Size = new System.Drawing.Size(238, 23);
            this.portTextBox.TabIndex = 23;
            this.portTextBox.Text = "587";
            // 
            // portLabel
            // 
            this.portLabel.AutoSize = true;
            this.portLabel.Font = new System.Drawing.Font("Microsoft Sans Serif", 11.25F, System.Drawing.FontStyle.Regular, System.Drawing.GraphicsUnit.Point, ((byte)(204)));
            this.portLabel.Location = new System.Drawing.Point(5, 66);
            this.portLabel.Name = "portLabel";
            this.portLabel.Size = new System.Drawing.Size(44, 18);
            this.portLabel.TabIndex = 22;
            this.portLabel.Text = "порт:";
            // 
            // passwordLabel
            // 
            this.passwordLabel.AutoSize = true;
            this.passwordLabel.Font = new System.Drawing.Font("Microsoft Sans Serif", 11.25F, System.Drawing.FontStyle.Regular, System.Drawing.GraphicsUnit.Point, ((byte)(204)));
            this.passwordLabel.Location = new System.Drawing.Point(5, 159);
            this.passwordLabel.Name = "passwordLabel";
            this.passwordLabel.Size = new System.Drawing.Size(62, 18);
            this.passwordLabel.TabIndex = 18;
            this.passwordLabel.Text = "пароль:";
            // 
            // passwordTextBox
            // 
            this.passwordTextBox.Font = new System.Drawing.Font("Microsoft Sans Serif", 10.2F, System.Drawing.FontStyle.Regular, System.Drawing.GraphicsUnit.Point, ((byte)(204)));
            this.passwordTextBox.Location = new System.Drawing.Point(5, 180);
            this.passwordTextBox.Name = "passwordTextBox";
            this.passwordTextBox.PasswordChar = '*';
            this.passwordTextBox.Size = new System.Drawing.Size(238, 23);
            this.passwordTextBox.TabIndex = 17;
            // 
            // hostTextBox
            // 
            this.hostTextBox.Font = new System.Drawing.Font("Microsoft Sans Serif", 10.2F, System.Drawing.FontStyle.Regular, System.Drawing.GraphicsUnit.Point, ((byte)(204)));
            this.hostTextBox.Location = new System.Drawing.Point(4, 41);
            this.hostTextBox.Margin = new System.Windows.Forms.Padding(2);
            this.hostTextBox.Name = "hostTextBox";
            this.hostTextBox.Size = new System.Drawing.Size(238, 23);
            this.hostTextBox.TabIndex = 19;
            this.hostTextBox.Text = "smtp.mail.ru";
            // 
            // hostLabel
            // 
            this.hostLabel.AutoSize = true;
            this.hostLabel.Font = new System.Drawing.Font("Microsoft Sans Serif", 11.25F, System.Drawing.FontStyle.Regular, System.Drawing.GraphicsUnit.Point, ((byte)(204)));
            this.hostLabel.Location = new System.Drawing.Point(5, 21);
            this.hostLabel.Name = "hostLabel";
            this.hostLabel.Size = new System.Drawing.Size(60, 18);
            this.hostLabel.TabIndex = 20;
            this.hostLabel.Text = "сервер:";
            // 
            // openFileDialog1
            // 
            this.openFileDialog1.FileName = "openFileDialog1";
            this.openFileDialog1.FileOk += new System.ComponentModel.CancelEventHandler(this.openFileDialog1_FileOk);
            // 
            // Form1
            // 
            this.AutoScaleDimensions = new System.Drawing.SizeF(6F, 13F);
            this.AutoScaleMode = System.Windows.Forms.AutoScaleMode.Font;
            this.ClientSize = new System.Drawing.Size(674, 402);
            this.Controls.Add(this.senderGroupBox);
            this.Controls.Add(this.recipientsGroupBox);
            this.Margin = new System.Windows.Forms.Padding(2);
            this.Name = "Form1";
            this.Text = "отправка сообщения";
            this.recipientsGroupBox.ResumeLayout(false);
            this.recipientsGroupBox.PerformLayout();
            this.senderGroupBox.ResumeLayout(false);
            this.senderGroupBox.PerformLayout();
            this.ResumeLayout(false);

        }

        #endregion

        private System.Windows.Forms.GroupBox recipientsGroupBox;
        private System.Windows.Forms.Label subjectLabel;
        private System.Windows.Forms.TextBox subjectTextBox;
        private System.Windows.Forms.RichTextBox bodyRichTextBox;
        private System.Windows.Forms.Label toLabel;
        private System.Windows.Forms.Label fromLabel;
        private System.Windows.Forms.TextBox toTextBox;
        private System.Windows.Forms.TextBox fromTextBox;
        private System.Windows.Forms.GroupBox senderGroupBox;
        private System.Windows.Forms.TextBox portTextBox;
        private System.Windows.Forms.Label portLabel;
        private System.Windows.Forms.Label passwordLabel;
        private System.Windows.Forms.TextBox passwordTextBox;
        private System.Windows.Forms.TextBox hostTextBox;
        private System.Windows.Forms.Label hostLabel;
        private System.Windows.Forms.Button attachButton;
        private System.Windows.Forms.TextBox attachmentTextBox;
        private System.Windows.Forms.Label attachLabel;
        private System.Windows.Forms.Button sendButton;
        private System.Windows.Forms.OpenFileDialog openFileDialog1;
    }
}