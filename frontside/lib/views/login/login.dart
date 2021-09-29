import 'package:flutter/material.dart';

class Login extends StatefulWidget {
  static const String route = "/";

  @override
  _Login createState() => _Login();
}

class _Login extends State<Login> {
  @override
  Widget build(BuildContext context) {
    var _whiteTextStyle = TextStyle(color: Colors.white);
    return Scaffold(
        body: SafeArea(
            child: Column(
                mainAxisAlignment: MainAxisAlignment.spaceBetween,
                children: <Widget>[
          Expanded(
              child: Container(
            color: Theme.of(context).colorScheme.primary,
            child: Padding(
              padding: const EdgeInsets.symmetric(horizontal: 14, vertical: 14),
              child: Column(
                children: [
                  TextFormField(
                    keyboardType: TextInputType.emailAddress,
                    style: _whiteTextStyle,
                    decoration: InputDecoration(
                      labelText: 'email or username',
                      border: UnderlineInputBorder(),
                      enabledBorder: UnderlineInputBorder(
                          borderSide: BorderSide(color: Colors.white)),
                      focusedBorder: UnderlineInputBorder(
                          borderSide: BorderSide(color: Colors.white)),
                      labelStyle: _whiteTextStyle,
                    ),
                  ),
                  SizedBox(
                    height: 16,
                  ),
                  Row(
                    children: [
                      Expanded(
                        child: SizedBox(),
                      ),
                      OutlinedButton(
                        onPressed: () => {},
                        child: Text("NEXT"),
                        style: ButtonStyle(
                            backgroundColor: MaterialStateProperty.all<Color>(
                                Theme.of(context).colorScheme.surface)),
                      ),
                    ],
                  )
                ],
              ),
            ),
          )),
          Column(
            children: [
              LoginWith(
                icon: Image(
                  image: AssetImage('assets/login/logo-google-G.png'),
                  fit: BoxFit.contain,
                  width: 18,
                ),
                title: "LOGIN WITH GOOGLE",
              ),
              LoginWith(
                icon: Image(
                  image: AssetImage('assets/login/logo-apple-black-100.png'),
                  fit: BoxFit.contain,
                  width: 18,
                ),
                title: "LOGIN WITH APPLE",
              ),
            ],
          )
        ])));
  }
}

class LoginWith extends StatelessWidget {
  const LoginWith({
    Key? key,
    required this.icon,
    required this.title,
  }) : super(key: key);

  final Widget icon;
  final String title;

  @override
  Widget build(BuildContext context) {
    return InkWell(
      onTap: () => {},
      child: Container(
          child: Padding(
        padding: const EdgeInsets.all(14),
        child: Row(children: [icon, SizedBox(width: 10), Text(title)]),
      )),
    );
  }
}
