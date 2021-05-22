import java.time.Duration;
import java.time.LocalDate;
import java.time.LocalDateTime;
import java.time.LocalTime;

public class Gigasecond {
    private static final Duration gigasecond = Duration.ofSeconds(1_000_000_000);
    private final LocalDateTime gigamoment;

    public Gigasecond(LocalDate moment) {
        this(LocalDateTime.of(moment, LocalTime.MIN));
    }

    public Gigasecond(LocalDateTime moment) {
        this.gigamoment = moment.plus(gigasecond);
    }

    public LocalDateTime getDateTime() {
        return this.gigamoment;
    }
}
