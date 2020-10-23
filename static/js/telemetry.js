var telemetry = new Vue({
    el: '#telemetry',
    data: {
        Status: {
            EngineDamage: 0,
            FrontLeftWingDamage: 0,
            FrontRightWingDamage: 0,
            GearBoxDamage: 0,
            RearWingDamage: 0,
            BrakesTemperature: [100, 100, 100, 100],
            EngineTemperature: 0,
            TyresInnerTemperature: [80, 80, 80, 80],
            TyresSurfaceTemperature: [80, 80, 80, 80],
            ErsDeployMode: 0,
            ErsStoreEnergy: 0,
            FuelMix: 0,
            FuelRemainingLaps: 0,
            TyresWear: [0, 0, 0, 0],
        },
        Session: {
            Weather: 0,
            TrackTemperature: 20,
            AirTemperature: 30,
            Position: 1,
            TotalCars: 1,
            LapsDone: 0,
            TotalLaps: 10,
            TrackLength: 4000,
            SessionType: "P1",
            Track: "",
            Country: "de",
            SessionTimeLeft: 0,
            SessionDuration: 0,
            WeatherForecasts: [],
        },
        Ranking: {
            NumCars: 1,
            RankingData: [{
                LapData: {
                    LastLapTime: 1.0,
                    BestLapTime: 1.0,
                    Diff: 1.0,
                    BestLapNum: 1,
                    CurrentLapNum: 1,
                    TotalDistance: 0,
                    DriverStatus: 0,
                    ResultStatus: 2,
                    BestLapSector1TimeInMS: 0,
                    BestLapSector2TimeInMS: 0,
                    BestLapSector3TimeInMS: 0,
                    BestOverallSector1TimeInMS: 0,
                    BestOverallSector2TimeInMS: 0,
                    BestOverallSector3TimeInMS: 0,
                    CarPosition: 1,
                    PitStatus: 0,
                    Penalties: 0,
                },
                IsPlayer: true,
                AI: 0,
                Team: "",
                TeamColor: "",
                Tyres: "dry",
                TyresBestLap: "",
                TyresAge: 0,
                Name: "",
                Nation: "",
                DrivenDistance: 0,
            }]
        },
        Result: {
            NumCars: 1,
            RaceResultData: [{
                Position: 1,
                GridPosition: 1,
                NumLaps: 1,
                Points: 25,
                Team: "",
                TeamColor: "",
                Name: "",
                Nation: "",
                NumPitStops: 0,
                ResultStatus: 0,
                BestLapTime: 0,
                TotalRaceTime: 0,
                TotalTimeDiff: 0,
                PenaltiesTime: 0,
                NumTyreStints: 0,
                Tyres: ["soft"]
            }]
        },
        Map: {
            Width: 100,
            Height: 100,
            TrackID: 0,
            XOffset: 0,
            ZOffset: 0,
            Positions: [{
                WorldPositionX: 0,
                WorldPositionZ: 0,
                TeamColor: "#000000",
                Name: "ABC"
            }]
        },
        Canvas: null,
        CanvasContext: null,
        CanvasProps: {
            Rotated: false,
            Translated: false,
        }
    },
    computed: {
        ERSCharge: function() {
            return Math.round(this.Status.ErsStoreEnergy / fullERS * 1000) / 10
        },
        RemainingLaps: function() {
            return Math.round(this.Status.FuelRemainingLaps * 100) / 100
        },
        BarWidth: function() {
            return "width:" + this.ERSCharge + "%"
        },
        BarColor: function() {
            if (this.ERSCharge > 90) {
                return "progress-bar bg-success"
            }
            if (this.ERSCharge > 50) {
                return "progress-bar bg-info"
            }
            if (this.ERSCharge > 20) {
                return "progress-bar bg-warning"
            }
            return "progress-bar bg-danger"
        },
        BrakesTemperatureAll: function() {
            return this.Status.BrakesTemperature[0] + "°C, " + this.Status.BrakesTemperature[1] + "°C, " + this.Status.BrakesTemperature[2] + "°C, " + this.Status.BrakesTemperature[3] + "°C"
        },
        TyresInnerTemperatureAll: function() {
            return this.Status.TyresInnerTemperature[0] + "°C, " + this.Status.TyresInnerTemperature[1] + "°C, " + this.Status.TyresInnerTemperature[2] + "°C, " + this.Status.TyresInnerTemperature[3] + "°C"
        },
        TyresSurfaceTemperatureAll: function() {
            return this.Status.TyresSurfaceTemperature[0] + "°C, " + this.Status.TyresSurfaceTemperature[1] + "°C, " + this.Status.TyresSurfaceTemperature[2] + "°C, " + this.Status.TyresSurfaceTemperature[3] + "°C"
        },
        TyresDamageAll: function() {
            return this.Status.TyresDamage[0] + "%, " + this.Status.TyresDamage[1] + "%, " + this.Status.TyresDamage[2] + "%, " + this.Status.TyresDamage[3] + "%"
        },
        ERSMode: function() {
            if (this.Status.ErsDeployMode == 1) {
                return "⚡"
            }
            if (this.Status.ErsDeployMode == 2) {
                return "⚡⚡⚡"
            }
            if (this.Status.ErsDeployMode == 3) {
                return "⚡⚡"
            }
            return "off"
        },
        WeatherString: function() {
            return 'wi ' + weatherIcons[this.Session.Weather]
        },
        CountryString: function() {
            return 'flag-icon flag-icon-' + this.Session.Country
        },
        SessionTimeLeftFormatted: function() {
            if (this.Session.SessionTimeLeft > 3599) {
                return moment.utc(moment.duration(this.Session.SessionTimeLeft, "seconds").asMilliseconds()).format("h:mm:ss")
            }
            return moment.utc(moment.duration(this.Session.SessionTimeLeft, "seconds").asMilliseconds()).format("mm:ss")
        },
        SessionDurationFormatted: function() {
            SessionTime = this.Session.SessionDuration - this.Session.SessionTimeLeft
            if (SessionTime > 3599) {
                return moment.utc(moment.duration(SessionTime, "seconds").asMilliseconds()).format("h:mm:ss")
            }
            return moment.utc(moment.duration(SessionTime, "seconds").asMilliseconds()).format("mm:ss")
        },
        FuelMixText: function() {
            if (this.Status.FuelMix == 1) {
                return "🔥🔥"
            }
            if (this.Status.FuelMix == 2) {
                return "🔥🔥🔥"
            }
            if (this.Status.FuelMix == 3) {
                return "🔥🔥🔥🔥"
            }
            return "🔥"
        }
    },
    methods: {
        TyreColor(i) {
            return "background-color:" + getColorForPercentage(1 - this.Status.TyresWear[i] / 100) + ";"
        },
        Weather(i) {
            return 'wi ' + weatherIcons[i]
        },
        Flag(country) {
            return 'flag-icon flag-icon-' + country
        },
        TyreCompoundImg(tyreCompound) {
            if (tyreCompound == "") {
                return ""
            }
            return "assets/" + tyreCompound + ".png"
        },
        TyreVisibility(tyreCompound) {
            if (tyreCompound == "") {
                return "visibility: hidden;"
            }
            return ""
        },
        TeamColorStyle(color) {
            return "background-color: #" + color + ";"
        },
        LapTimeFormatted(lapTime) {
            if (lapTime > 3599) {
                return moment.utc(moment.duration(lapTime, "seconds").asMilliseconds()).format("h:mm:ss.SSS")
            } else {
                return moment.utc(moment.duration(lapTime, "seconds").asMilliseconds()).format("mm:ss.SSS")
            }
        },
        SectorTimeFormatted(sectorTime) {
            return moment.utc(moment.duration(sectorTime, "milliseconds").asMilliseconds()).format("ss.SSS")
        },
        RowBackground(IsPlayer, AI) {
            if (IsPlayer) {
                return "background: #0a6e17;"
            }
            if (!AI) {
                return "background: #1d3f96;"
            }
            return ""
        },
        DriverStatusIcon(i) {
            switch (i) {
                case 0:
                    return "▲"
                case 1:
                    return "🏁"
                default:
                    return "●"
            }
        },
        DriverStatusStyle(i) {
            switch (i) {
                case 0:
                    return "text-align:center; color: #e8e84f;"
                case 1:
                    return "text-align:center;"
                case 2:
                    return "text-align:center; color: #c71616;"
                case 3:
                    return "text-align:center; color: #66e649;"
                case 4:
                    return "text-align:center; color: #c71616;"
                default:
                    return ""
            }
        },
        PositionOrResult(status, pos) {
            if (status == 6 || status == 7) {
                return "DNF"
            }
            if (status == 4) {
                return "DSQ"
            }
            return pos
        },
        GridPosDiff(pos, grid) {
            if ((pos - grid) > 0) {
                return '▼'
            } else if ((pos - grid < 0)) {
                return '▲'
            } else {
                return "="
            }
        },
        GridPosDiffStyle(pos, grid) {
            if ((pos - grid) > 0) {
                return "text-align: center; color: #c71616;"
            } else if ((pos - grid < 0)) {
                return "text-align: center; color: #4fb837;"
            } else {
                return "text-align: center;"
            }
        }
    },
    watch: {
        Map: function(val) {
            if (this.Canvas.getAttribute("height") != val.Height) {
                this.Canvas.setAttribute("height", val.Height);
            }
            if (this.Canvas.getAttribute("width") != val.Width) {
                this.Canvas.setAttribute("width", val.Width);
            }
            if (this.CanvasProps.Translated == false) {
                this.CanvasContext.rotate(val.Rotation * Math.PI / 180);
                this.CanvasContext.translate(val.XTranslation, val.YTranslation);
                this.CanvasProps.Translated = true;
            }

            this.CanvasContext.strokeStyle = "#000000";
            this.CanvasContext.lineWidth = 2;
            this.CanvasContext.clearRect(-10000, -10000, 100000, 100000);
            for (i = 0; i < 20; i++) {
                var x = val.Positions[i].WorldPositionX;
                var z = val.Positions[i].WorldPositionZ;
                var color = val.Positions[i].TeamColor;
                var name = val.Positions[i].Name;
                // Un-Rotate to sign is upright
                this.CanvasContext.translate(x, z);
                this.CanvasContext.rotate(-val.Rotation * Math.PI / 180);
                this.CanvasContext.font = "40px monospace"
                this.CanvasContext.fillStyle = "#212121";
                // Draw sign above circle
                this.CanvasContext.fillRect(-42, -70, 80, 50)
                this.CanvasContext.strokeRect(-42, -70, 80, 50)
                this.CanvasContext.fillStyle = "#" + color;
                this.CanvasContext.fillText(name, -35, -30)
                this.CanvasContext.rotate(val.Rotation * Math.PI / 180);
                this.CanvasContext.translate(-x, -z);
                // Draw outer black circle
                this.CanvasContext.beginPath();
                this.CanvasContext.arc(x, z, 17, 0, 2 * Math.PI);
                this.CanvasContext.stroke();
                // Draw inner filled circle
                this.CanvasContext.beginPath();
                this.CanvasContext.arc(x, z, 16, 0, 2 * Math.PI);
                this.CanvasContext.fill();
            }
        }
    },
    mounted: function() {
        this.Canvas = document.getElementById("map");
        var ctx = this.Canvas.getContext("2d");
        this.CanvasContext = ctx;
    }
});