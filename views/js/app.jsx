class App extends React.Component {
    constructor(props) {
        super(props);
        this.state = {"loggedIn": false};
    }

    serverRequest() {
        $.get("http://localhost:3000/api/login-status", res => {
            if (res.status == true) {
                this.setState({
                    "loggedIn": true,
                })
            }
        })
    }

    componentDidMount() {
        this.serverRequest();
    }

    render() {
        if (this.state.loggedIn) {
            return (
                <LoggedIn />
            );
        } else {
            return (
                <Home />
            )
        }
    }
}

class Home extends React.Component {
    render() {
        return (
            <div className="container">
                <div className="col-xs-8 col-xs-offset-2 jumbotron text-center">
                    <h1> Crypto Watch Test Socket </h1>
                    <a href="/login">SignIn</a>
                </div>
            </div>
        )
    }
}

class LoggedIn extends React.Component {
    constructor(props) {
        super(props)
        this.state = {
            "ethPrice": 0,
            "updatedAt": 0,
        };
    }

    componentDidMount() {
        const watchUpdateInterval = setInterval(() => {
            $.get("http://localhost:3000/api/watch", res => {
                if ((res.marketUpdate.tradesUpdate.trades != null) && (res.marketUpdate.tradesUpdate.trades.length > 0)) {
                    this.setState({
                        "ethPrice": res.marketUpdate.tradesUpdate.trades[0].priceStr,
                        "updatedAt": res.marketUpdate.tradesUpdate.trades[0].timestamp,
                    })
                }
            })
        }, 1000);
        return () => clearInterval(watchUpdateInterval);
    }

    render() {
        return (
            <div className="container">
                <div className="col-lg-12">
                    <br />
                    <span className="pull-right">
                        <a href="/logout">LogOut</a>
                    </span>
                    <h2>
                        This is the home page
                    </h2>
                    <div>
                        Current Price of Ether: {this.state.ethPrice}
                    </div>
                </div>
            </div>
        )
    }
}

ReactDOM.render(<App />, document.getElementById('app'));