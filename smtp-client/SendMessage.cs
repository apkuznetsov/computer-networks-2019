using System.Net;
using System.Net.Mail;
using System.Net.Mime;

namespace Lab4
{
    public static class SendMessage
    {
        public static void Send(
            string from_,
            string to_,
            string subject_,
            string body_,
            string attachment_,
            string host_,
            int port_,
            string password_)
        {
            #region mail
            MailMessage mailMessage = new MailMessage();
            mailMessage.From = new MailAddress(from_);  // From: адрес отправителя. Представляет объект MailAddress.

            to_.Replace(" ", "");
            string[] toArr = to_.Split(',');
            foreach (string recipient in toArr)
                mailMessage.To.Add(recipient);      // To: адрес получателя. Также представляет объект MailAddress.

            mailMessage.Subject = subject_;  // Subject: определяет тему письма.
            mailMessage.Body = body_;    // Body: непосредственно текст письма.

            if (attachment_ != "")
            {
                Attachment attachment = new Attachment(attachment_, MediaTypeNames.Application.Pdf);
                mailMessage.Attachments.Add(attachment);
            }
            #endregion

            #region client
            /* Для отправки почты в среде интернет используется протокол SMTP 
             * (Simple Mail Transfer Protocol). 
             * Данный протокол указывает, 
             * как почтовые сервера взаимодействуют при передаче электронной почты.
             * 
             * Host:        SMTP-сервер, с которого производится отправление почты.
             * Port:        порт, используемый SMP-сервером.
             * Credentials: аутентификационные данные отправителя. */
            SmtpClient smtpClient = new SmtpClient();

            smtpClient.Host = host_;
            smtpClient.Port = port_;

            string userName = from_.Substring(0, from_.IndexOf(@"@"));
            smtpClient.Credentials = new NetworkCredential(userName, password_);

            /* EnableSsl: указывает,
             * будет ли использоваться протокол SSL при отправке. */
            smtpClient.EnableSsl = true;

            smtpClient.DeliveryMethod = SmtpDeliveryMethod.Network; // Электронная почта отправляется по сети на сервер SMTP.
            #endregion

            #region sending
            smtpClient.Send(mailMessage);

            smtpClient.Dispose();
            mailMessage.Dispose();
            #endregion
        }
    }
}