const fullERS = 4000000;
const weatherIcons = ['wi-day-sunny', 'wi-day-cloudy', 'wi-cloud', 'wi-sprinkle', 'wi-rain', 'wi-thunderstorm']

function hideOverlay() {
    document.getElementById("overlay").style.display = "none";
}

var percentColors = [
    { pct: 0.0, color: { r: 0xff, g: 0x00, b: 0 } },
    { pct: 0.5, color: { r: 0xff, g: 0xff, b: 0 } },
    { pct: 1.0, color: { r: 0x00, g: 0xff, b: 0 } }
];

function getColorForPercentage(pct) {
    for (var i = 1; i < percentColors.length - 1; i++) {
        if (pct < percentColors[i].pct) {
            break;
        }
    }
    var lower = percentColors[i - 1];
    var upper = percentColors[i];
    var range = upper.pct - lower.pct;
    var rangePct = (pct - lower.pct) / range;
    var pctLower = 1 - rangePct;
    var pctUpper = rangePct;
    var color = {
        r: Math.floor(lower.color.r * pctLower + upper.color.r * pctUpper),
        g: Math.floor(lower.color.g * pctLower + upper.color.g * pctUpper),
        b: Math.floor(lower.color.b * pctLower + upper.color.b * pctUpper)
    };
    return 'rgb(' + [color.r, color.g, color.b].join(',') + ',0.4)';
}

function Sleep(milliseconds) {
    return new Promise(resolve => setTimeout(resolve, milliseconds));
}

function connect() {
    var conn = new ReconnectingWebSocket("ws://localhost:8080/ws");
    conn.onclose = function(evt) {
        console.log('Socket is closed. Reconnect will be attempted.');
        telemetry.Connected = false;
    }
    conn.onopen = function() {
        telemetry.Connected = true;
    }
    conn.onmessage = function(evt) {
        messages = evt.data.split("----");
        for (message of messages) {
            try {
                data = JSON.parse(message);
            } catch (err) {
                console.log("too much data!"); // happens when 2 messages are sent together
                return
            }

            //console.log(data);
            switch (data.MessageType) {
                case "SessionInfo":
                    telemetry.Session = data;
                    break;
                case "RankingInfo":
                    data.RankingData = data.RankingData.slice(0, data.NumCars)
                    for (i = 1; i < data.NumCars; i++) {
                        if (data.RankingData[i].LapData.BestLapTime == 0 && data.SessionType != 'R') {
                            continue
                        }
                        data.RankingData[i].LapData.Diff = (data.RankingData[i].LapData.BestLapTime - data.RankingData[i - 1].LapData.BestLapTime) * 1000
                        data.RankingData[i].LapData.DistDiff = (Math.round(data.RankingData[i - 1].LapData.TotalDistance - data.RankingData[i].LapData.TotalDistance) / 1000).toFixed(3)
                    }
                    telemetry.Ranking = data;
                    break;
                case "CarStatusInfo":
                    telemetry.Status = data;
                    break;
                case "RaceResult":
                    data.RaceResultData = data.RaceResultData.slice(0, data.NumCars)
                    for (i = 0; i < data.NumCars; i++) {
                        data.RaceResultData[i].Tyres = data.RaceResultData[i].Tyres.slice(0, data.RaceResultData[i].NumTyreStints)
                        if (i == 0) {
                            continue
                        }
                        if (data.RaceResultData[i].ResultStatus == 3) {
                            data.RaceResultData[i].TotalTimeDiff = data.RaceResultData[i].TotalRaceTime - data.RaceResultData[i - 1].TotalRaceTime
                        }
                    }
                    telemetry.Result = data;
                    document.getElementById("overlay").style.display = "block";
                    break;
                case "PositionInfo":
                    data.Positions = data.Positions.slice(0, data.NumCars)
                    telemetry.Map = data;
                    break;
                default:
                    break;
            }
        }
    }
}
