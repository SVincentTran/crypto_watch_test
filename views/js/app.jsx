let userId = ""

class App extends React.Component {
    constructor(props) {
        super(props);
        this.state = {"loggedIn": false};
    }

    serverRequest() {
        $.get(url + "/api/login-status", res => {
            if (res.status == true) {
                userId = res.user_id
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
                    <a href="/login" className="btn btn-primary btn-lg">Sign In</a>
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
            "userId": "",
        };
    }

    componentDidMount() {
        $.get(url + "/api/login-status")

        const watchUpdateInterval = setInterval(() => {
            $.get(url + "/api/watch", res => {
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
                        <a href="/logout" className="btn btn-primary">Log Out</a>
                    </span>
                    <h2>
                        Crypto Dashboard
                    </h2>
                </div>
                <div className="col-lg-12">
                    <div className="container text-center">
                        Current Price of Ethereum:
                        <mark>
                            <strong>
                            {this.state.ethPrice}
                            </strong>
                        </mark>
                    </div>
                    <br />
                    <br />
                    <br />
                    <hr />
                </div>

                <div className="col-lg-12">
                    <div className="container text-center">
                        <p> UserId: {userId} </p>
                    </div>
                </div>
            </div>
        )
    }
}

ReactDOM.render(<App />, document.getElementById('app'));