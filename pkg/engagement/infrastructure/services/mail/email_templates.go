package mail

// MarketingEmailTemplate generates an simple marketing email template
const MarketingEmailTemplate = `
<!DOCTYPE html>
<html>

<head>
    <title>Be.Well by Slade360° - Simple. Caring. Trusted. </title>
    <meta property="description" content="Join Be.Well today to take charge of your benefits">
    <!--VIEWPORT-->
    <meta name="viewport" content="width=device-width; initial-scale=1.0; maximum-scale=1.0; user-scalable=no;">
    <meta name="viewport" content="width=600, initial-scale = 2.3, user-scalable=no">
    <meta name="viewport" content="width=device-width">
    <!--CHARSET-->
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8" />

    <!-- IE=Edge and IE=X -->
    <meta http-equiv="X-UA-Compatible" content="IE=7" />
    <meta http-equiv="X-UA-Compatible" content="IE=8" />
    <meta http-equiv="X-UA-Compatible" content="IE=9" />
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <!--[if !mso]>-->
    <meta http-equiv="X-UA-Compatible" content="IE=edge" />
    <!--<![endif]-->
    <!-- INLINE STYLES -->
    <style type="text/css">
        @media screen {
            @font-face {
                font-family: "Lato";
                font-style: normal;
                font-weight: 400;
                src: local("Lato Regular"), local("Lato-Regular"),
                    url(https://fonts.gstatic.com/s/lato/v11/qIIYRU-oROkIk8vfvxw6QvesZW2xOQ-xsNqO47m55DA.woff) format("woff");
            }

            @font-face {
                font-family: "Lato";
                font-style: normal;
                font-weight: 700;
                src: local("Lato Bold"), local("Lato-Bold"),
                    url(https://fonts.gstatic.com/s/lato/v11/qdgUG4U09HnJwhYI-uK18wLUuEpTyoUstqEm5AMlJo4.woff) format("woff");
            }

            @font-face {
                font-family: "Lato";
                font-style: italic;
                font-weight: 700;
                src: local("Lato Bold Italic"), local("Lato-BoldItalic"),
                    url(https://fonts.gstatic.com/s/lato/v11/HkF_qI1x_noxlxhrhMQYELO3LdcAZYWl9Si6vvxL-qU.woff) format("woff");
            }
        }

        body,
        table,
        td {
            -webkit-text-size-adjust: 100%;
            -ms-text-size-adjust: 100%;
        }

        img {
            -ms-interpolation-mode: bicubic;
        }

        img {
            border: 0;
            height: auto;
            line-height: 100%;
            outline: none;
            text-decoration: none;
        }

        table {
            border-collapse: collapse !important;
        }

        body {
            height: 100% !important;
            margin: 0 !important;
            padding: 0 !important;
            width: 100% !important;
        }

        /* MOBILE STYLES */
        @media screen and (max-width: 600px) {
            h1 {
                font-size: 32px !important;
                line-height: 32px !important;
            }
        }

        /* ANDROID CENTER FIX */
        div[style*="margin: 16px 0;"] {
            margin: 0 !important;
        }
    </style>
</head>

<body style="
      background-color: #f4f4f4;
      margin: 0 !important;
      padding: 0 !important;
    ">
    <table border="0" cellpadding="0" cellspacing="0" width="100%" height="100%">
        <tr>
            <td bgcolor="#7B54C4" align="center">
                <table border="0" cellpadding="0" cellspacing="0" width="100%" style="max-width: 600px">
                    <tr>
                        <td align="center" valign="top" style="padding: 40px 10px 40px 10px"></td>
                    </tr>
                </table>
            </td>
        </tr>
        <tr>
            <td bgcolor="#7B54C4" align="center" style="padding: 0px 10px 0px 10px">
                <table border="0" cellpadding="0" cellspacing="0" width="85%" style="max-width: 600px">
                    <tr>
                        <td bgcolor="#ffffff" valign="top" style="
                  padding: 40px 20px 10px 24px;
                  border-radius: 4px 4px 0px 0px;
                  color: #111111;
                  font-family: 'Lato', Helvetica, Arial, sans-serif;
                  font-size: 48px;
                  font-weight: 400;
                  line-height: 48px;
                ">
                            <img src="https://lh3.googleusercontent.com/pw/ACtC-3fN_p8U8EZgmtQymnwrhr_-5Go6Kw5e5U7lkjyk1jjMIEwSs6rDNELplpgVk2IciMfw5AbnphxJYwdocnsE6Y88xyKGlNXm1E1x3Sm9uxeMHhwjf8YgNwo622G8cb-d7ntlbNl7-uPCEylu5O_KzZY=s638-no"
                                width="125" height="120" style="display: block; border: 0px; margin-bottom: 0" />
                        </td>
                    </tr>
                </table>
            </td>
        </tr>
        <tr>
            <td bgcolor="#f4f4f4" align="center" style="padding: 0px 10px 0px 10px">
                <table border="0" cellpadding="0" cellspacing="0" width="85%" style="max-width: 600px">
                    <tr>
                        <td bgcolor="#ffffff" align="left" style="
                  padding: 20px 30px 20px 30px;
                  color: #666666;
                  font-family: 'Lato', Helvetica, Arial, sans-serif;
                  font-size: 18px;
                  font-weight: 400;
                  line-height: 25px;
                ">
                            <p style="margin: 0;font-size: 25px;font-weight: bold; color: black;">Join Be.Well today to
                                take charge of
                                your benefits</p>
                            <p></p>

                            <p style="margin-bottom: 2px;">Hi,</p>
                            <p style="margin: 0">
                                Click a link below to download the new Be.Well App
                                on your <a
                                    href="https://play.google.com/store/apps/details?id=com.savannah.bewell">Android</a>
                                or <a href="https://apps.apple.com/ke/app/be-well-by-slade360/id1496576692">iOS</a>
                                device.
                            </p>
                            <br>
                            <p style="margin-bottom: 4px;font-size: 16px;font-weight: bold; color: black;">The Be.Well
                                app lets you:
                            </p>
                            <ul style="margin: 0;">
                                <li>Access your virtual Slade ID Wellness Card</li>
                                <li>Manage your benefits</li>
                                <li>View your claims</li>
                                <li>See your preauthorizations</li>
                                <li>Learn more about your health</li>
                                <li>Review your children's claims</li>
                                <li>See where your cover has been used</li>
                            </ul>

                            <p style="font-size: medium; margin-bottom: 0;">Click a link below to download the new
                                Be.Well app.</p>
                            <p style="display: flex;">
                                <a target="_blank"
                                    href="https://play.google.com/store/apps/details?id=com.savannah.bewell"
                                    style="margin-right: 20px;">
                                    <img src="https://a.bewell.co.ke/image/playstore.png" width="150"
                                        style="display: block; border: 0px; margin-bottom: 0" />
                                </a>

                                <a target="_blank"
                                    href="https://apps.apple.com/ke/app/be-well-by-slade360/id1496576692">
                                    <img src="
                    https://a.bewell.co.ke/image/appstore.png" width="150"
                                        style="display: block; border: 0px; margin-bottom: 0" />
                                </a>
                            </p>
                            <br>
                            <p style="margin: 0; ">
                                Thank you for using Be.Well.
                            </p>
                        </td>
                    </tr>

                    <tr>
                        <td bgcolor="#ffffff" align="center" style="
                  color: #666666;
                  font-family: 'Lato', Helvetica, Arial, sans-serif;
                  font-size: 40px;
                  font-weight: 900;
                  line-height: 40px;
                "></td>
                    </tr>

                    <tr>
                        <td bgcolor="#ffffff" align="left" style="
                  padding: 0px 30px 40px 30px;
                  border-radius: 0px 0px 4px 4px;
                  color: #666666;
                  font-family: 'Lato', Helvetica, Arial, sans-serif;
                  font-size: 18px;
                  font-weight: 400;
                  line-height: 25px;
                ">
                            <p style="margin: 0">
                                Regards,<br />
                                <span style="color: black; font-weight: 600;">Kevin from Be.Well</span>
                            </p>
                        </td>
                    </tr>
                </table>
            </td>
        </tr>
        <table border="0" cellpadding="0" cellspacing="0" width="100%" style="padding-top: 40px">
            <tr>
                <td bgcolor="#f4f4f4" align="center" style="
              padding: 40px 30px 40px 30px;
              border-radius: 0px 0px 4px 4px;
              color: #000000;
              font-family: 'Lato', Helvetica, Arial, sans-serif;
              font-size: 18px;
              font-weight: 400;
              line-height: 25px;
            ">
                    <p style="margin: 0">
                        For more information or queries, contact us at
                        <a href="mailto:feedback@bewell.co.ke">feedback@bewell.co.ke</a>
                        <br>
                        or call <a href="tel:0790360360">0790 360 360</a>
                    </p>
                </td>
            </tr>
        </table>
    </table>
    <script src="https://cdn.jsdelivr.net/npm/publicalbum@latest/embed-ui.min.js" async></script>
    <!-- Start of HubSpot Embed Code -->
    <script type="text/javascript" id="hs-script-loader" async defer src="//js.hs-scripts.com/20198195.js"></script>
    <!-- End of HubSpot Embed Code -->
    <!-- The core Firebase JS SDK is always required and must be listed first -->
    <script src="https://www.gstatic.com/firebasejs/8.7.1/firebase-app.js"></script>

    <script src="https://www.gstatic.com/firebasejs/8.7.1/firebase-analytics.js"></script>

    <script>
        var firebaseConfig = {
            apiKey: "AIzaSyAv2aRsSSHkOR6xGwwaw6-UTkvED3RNlBQ",
            authDomain: "bewell-app.firebaseapp.com",
            databaseURL: "https://bewell-app.firebaseio.com",
            projectId: "bewell-app",
            storageBucket: "bewell-app.appspot.com",
            messagingSenderId: "841947754847",
            appId: "1:841947754847:web:6304157d32c82fd96686ea",
            measurementId: "G-6XTZEB5070"
        };
        firebase.initializeApp(firebaseConfig);
        const analytics = firebase.analytics();


        analytics.logEvent('opened_join_bewell_marketing_email');
    </script>
</body>

</html>
`
