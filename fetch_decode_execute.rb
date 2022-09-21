class ClockSpeed
    def self.call
        empty_loop
        loop_with_math
    end

    def self.empty_loop
        rounds = 1000000000
        start_time = Time.now
        rounds.times do
            # do nothing
        end
        end_time = Time.now
        elapsed_time = end_time - start_time

        puts (rounds / elapsed_time)
    end

    def self.loop_with_math
        rounds = 1000000000
        num = 0
        start_time = Time.now
        rounds.times do
            num * 2
        end
        end_time = Time.now
        elapsed_time = end_time - start_time
        
        puts (rounds / elapsed_time)
    end
end