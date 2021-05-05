import java.time.Duration;
import java.time.LocalDate;
import java.time.LocalDateTime;
import java.time.LocalTime;

public class Gigasecond {
    private final LocalDateTime gigamoment;
    private final Duration gigasecond = Duration.ofSeconds(1_000_000_000);

    public Gigasecond(LocalDate moment) {
        this.gigamoment = LocalDateTime.of(moment, LocalTime.MIN).plus(gigasecond);
    }

    public Gigasecond(LocalDateTime moment) {
        this.gigamoment = moment.plus(gigasecond);
    }

    public LocalDateTime getDateTime() {
        return this.gigamoment;
    }
}
