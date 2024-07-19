# musegen

## Explore the Depths of Musical Structure

`musegen` is an innovative, open-source tool designed to help musicians, composers, and music enthusiasts explore the vast landscape of musical structures. By generating chord progressions across a wide array of genres, `musegen` opens up new avenues for creativity and learning in music composition.

### 🎵 Features

- **Multi-Genre Support**: From classical to EDM, blues to salsa, `musegen` covers a broad spectrum of musical styles.
- **Theoretically Grounded**: Each progression is rooted in general music theory, providing a solid foundation for learning and experimentation.
- **Extensible**: Easily add new genres or modify existing ones to suit your needs.
- **Exploratory by Nature**: Discover new chord combinations and progressions you might not have considered.

### 🚀 Getting Started

To begin your journey with `musegen`, simply clone the repository and build the project:

```bash
git clone https://github.com/yourusername/musegen.git
cd musegen
go build
```

Run `musegen` to start generating progressions:

```bash
$ musegen

A CLI tool to generate music-related content.

Usage:
  musegen [command]

Available Commands:
  completion        Generate the autocompletion script for the specified shell
  gen-chords        chord progression generator
  help              Help about any command
  list-config       chord progression generator config options
  list-progressions list built-in progressions

Flags:
  -h, --help   help for musegen

Use "musegen [command] --help" for more information about a command.
```

### 🎹 Usage

`musegen` offers a simple command-line interface for generating progressions:

```bash
$ musegen gen-chords

generate chord progressions based on specified key, scale, and genre.

Usage:
  musegen gen-chords [flags]

Flags:
  -a, --allow-weird-shit   allow the generator to break some default rules
  -g, --genre string       genre of the progression (e.g., Pop, Jazz, Blues) (default "Rock")
  -h, --help               help for gen-chords
  -k, --key string         key of the progression (e.g., C, F#, Bb) (default "C")
  -r, --random             generate a completely random progression (1760 possibilities)
  -s, --scale string       scale type (e.g., Major, NaturalMinor, Dorian) (default "Major")
```

Explore different genres and progression lengths to find inspiration for your next composition!

### 🧠 The Philosophy Behind `musegen`

Music creation is a journey of exploration and discovery. `musegen` is designed not to replace human creativity, but to enhance it. By providing a tool for rapid generation and exploration of chord progressions, we hope to inspire musicians to push the boundaries of their compositions and discover new musical territories.

Whether you're a seasoned composer looking for fresh ideas, a music theory student seeking to understand progression structures, or a curious musician wanting to explore different genres, `musegen` is your companion in the vast world of musical possibilities.

### 🤝 Contributing

We welcome contributions from the community! Whether it's adding new genres, improving existing algorithms, or enhancing documentation, your input is valuable. Please see our [CONTRIBUTING.md](CONTRIBUTING.md) file for guidelines on how to contribute.

### 📜 License

`musegen` is licensed under the GNU General Public License v3.0 (GPLv3). This means you're free to use, modify, and distribute the software, but any derivatives must also be open source under the same license. For more details, see the [LICENSE](LICENSE) file.

### 🙏 Acknowledgements

`musegen` stands on the shoulders of giants. We're grateful to the countless musicians, theorists, and composers whose work has informed our understanding of musical structures and made this tool possible.

### 🧠 Made with AI

`musegen`'s progressions were generated with [Claude 3.5 Sonnet](https://www.anthropic.com/news/claude-3-5-sonnet). They are not perfect, will likely contain various errors, and may not be suitable for all musical contexts. Use them as a starting point for your own compositions and adapt them to your needs. If you see fundamental issues with the progressions, please let me know by opening an issue on GitHub or submitting a pull request.

### 🎶 Builtin Progressions

These are the current set of progressions that `musegen` can generate:

<details><summary>Expand</summary>
```shell
$ musegen list-progressions

Genre             Progression Name                                                           Chord Progression In C Major
-----             ----------------                                                           ----------------------------
Blues             12-Bar Blues (Basic)                                                       C7 - C7 - C7 - C7 - D#7 - D#7 - C7 - C7 - E7 - D#7 - C7 - E7
Blues             12-Bar Blues (Quick Change)                                                C7 - D#7 - C7 - C7 - D#7 - D#7 - C7 - C7 - E7 - D#7 - C7 - E7
Blues             12-Bar Blues (Minor)                                                       Cmin - Cmin - Cmin - Cmin - D#min - D#min - Cmin - Cmin - E7 - D#7 - Cmin - Cmin
Blues             8-Bar Blues                                                                C7 - C7 - D#7 - D#7 - E7 - D#7 - C7 - C7
Blues             16-Bar Blues                                                               C7 - C7 - C7 - C7 - D#7 - D#7 - C7 - C7 - E7 - E7 - C7 - C7 - E7 - D#7 - C7 - C7
Blues             Stormy Monday Blues                                                        C7 - D#7 - Fmin7 - C#min7 - D#7 - G#7 - C7 - G#7 - D#7 - F7 - Cmaj7 - E7
Blues             St. Louis Blues                                                            C7 - D#7 - F#7 - F7 - C7 - E7 - C7 - E7
Blues             Jazz Blues                                                                 C7 - D#7 - F#7 - F7 - D#7 - G#7 - C7 - G#7 - D#7 - F7 - C7 - E7
Blues             Boogie Woogie                                                              C7 - C7 - C7 - C7 - D#7 - D#7 - C7 - C7 - E7 - D#7 - C7 - C7
Bossa Nova        ii-V-I (Basic Bossa Nova)                                                  C#min7 - E7 - Cmaj7
Bossa Nova        i-IV7-iii7-bVI7-ii7-V7 (Minor Bossa Nova)                                  Cmin7 - D#7 - Dmin7 - F7 - C#min7 - E7
Bossa Nova        I-vi-ii-V (Bossa Nova Turnaround)                                          Cmaj7 - Fmin7 - C#min7 - E7
Bossa Nova        Imaj7-vi7-iimin7-V7 (Extended Bossa Nova)                                  Cmaj7 - Fmin7 - C#min7 - E7
Bossa Nova        I-#Idim7-ii7-V7 (Chromatic Bossa Nova)                                     Cmaj7 - C#dim7 - C#min7 - E7
Bossa Nova        Imaj7-bVIImaj7-bIIImaj7-bVImaj7 (Jobim - Corcovado)                        Cmaj7 - F#maj7 - Dmaj7 - Fmaj7
Bossa Nova        Imaj7-#Idim7-ii7-V7 (Jobim - The Girl from Ipanema)                        Cmaj7 - C#dim7 - C#min7 - E7
Bossa Nova        i7-IV7-iii7-VI7 (João Gilberto - Desafinado)                               Cmin7 - D#7 - Dmin7 - F7
Bossa Nova        i-bVII-bVI-V7 (Luiz Bonfá - Manhã de Carnaval)                             Cmin - F#maj - Fmaj - E7
Bossa Nova        Imaj7-vi7-iimin7-V7-iii7-VI7-ii7-V7 (Jobim - Wave)                         Cmaj7 - Fmin7 - C#min7 - E7 - Dmin7 - F7 - C#min7 - E7
Bossa Nova        i-V7/III-IIImaj7-V7/VI-VImaj7-iv7-V7 (Jobim - How Insensitive)             Cmin - E7 - Dmaj7 - G7 - Fmaj7 - D#min7 - E7
Bossa Nova        Imaj7-vi7-ii7-V7-iii7-VI7-ii7-V7 (João Donato - A Rã)                      Cmaj7 - Fmin7 - C#min7 - E7 - Dmin7 - F7 - C#min7 - E7
Classical         Perfect Authentic Cadence                                                  E - C
Classical         Plagal Cadence                                                             D# - C
Classical         Deceptive Cadence                                                          E - Fmin
Classical         Half Cadence                                                               C - E
Classical         Pachelbel's Canon                                                          C - E - F - Dmin - D# - C - D# - E
Classical         Circle of Fifths (Partial)                                                 C - D# - G# - C# - F# - B - E - C
Classical         Romanesca                                                                  C - E - Fmin - Dmin - D# - C - D# - E
Classical         La Folia                                                                   C#min - F - C - D# - C#min - F - C#min
Classical         Mozart's Sonata Facile (Opening)                                           C - E - C - F - C
Classical         Beethoven's Für Elise (Opening)                                            Cmin - E - Cmin - E - Cmin
Classical         Bach's Prelude in C Major (Opening)                                        C - C#min - E - C - D# - C#min - E - C
Country           I-IV-V (Basic Country)                                                     C - D# - E
Country           I-V-IV-I (Country Turnaround)                                              C - E - D# - C
Country           I-IV-I-V (Nashville Progression)                                           C - D# - C - E
Country           vi-IV-I-V (Modern Country)                                                 Fmin - D# - C - E
Country           I-V-vi-IV (Pop Country)                                                    C - E - Fmin - D#
Country           I-IV-V-IV (Country Walk-up)                                                C - D# - E - D#
Country           I-III-IV-iv (Country Blues)                                                C - D - D# - D#min
Country           I-vi-ii-V (Country Jazz)                                                   C - Fmin - C#min - E
Country           I-V-vi-iii-IV-I-IV-V (Extended Country)                                    C - E - Fmin - Dmin - D# - C - D# - E
Country           I-bVII-IV (Country Rock)                                                   C - F# - D#
Country           I-ii-V-I (Country Ballad)                                                  C - C#min - E - C
Country           I-IV-V-V7 (Country Shuffle)                                                C - D# - E - E7
Disco             I-V-vi-IV (Basic Disco)                                                    Cmaj7 - E7 - Fmin7 - D#maj7
Disco             i-iv-V-i (Minor Disco)                                                     Cmin7 - D#min7 - E7 - Cmin7
Disco             I-vi-ii-V (Jazz-influenced Disco)                                          Cmaj9 - Fmin9 - C#min9 - E9
Disco             I-bVII-IV-V (Funk-inspired Disco)                                          C9 - F#9 - D#9 - E9
Disco             Im7-bVII7-bVI7-V7 (Chromatic Disco)                                        Cmin7 - F#7 - F7 - E7
Disco             I-IV-V-IV (Chic - Le Freak)                                                C9 - D#9 - E9 - D#9
Disco             i-bVII-bVI-V (Donna Summer - I Feel Love)                                  Cmin - F# - F - E
Disco             I-V-vi-IV-I-ii-V (ABBA - Dancing Queen)                                    Cmaj - E - Fmin - D# - Cmaj - C#min - E
Disco             I-IV-V (Sister Sledge - We Are Family)                                     C7 - D#7 - E7
Disco             Im7-IV7-bVII7 (Bee Gees - Stayin' Alive)                                   Cmin7 - D#7 - F#7
Disco             I-bVII-IV (KC and the Sunshine Band - That's the Way (I Like It))          C7 - F#7 - D#7
Disco             i-IV-V-i (Gloria Gaynor - I Will Survive)                                  Cmin - D#maj - Emaj - Cmin
Disco             I-V-vi-iii-IV-I-IV-V (Earth, Wind & Fire - September)                      Cmaj9 - E9 - Fmin9 - Dmin9 - D#maj9 - Cmaj9 - D#maj9 - E9
EDM               I-V-vi-IV (Progressive House)                                              C - E - Fmin - D#
EDM               vi-IV-I-V (Trance)                                                         Fmin - D# - C - E
EDM               I-V-vi-iii-IV (Melodic Dubstep)                                            C - E - Fmin - Dmin - D#
EDM               i-VI-III-VII (Deep House)                                                  Cmin - F - D - F#
EDM               I-iii-vi-IV (Future Bass)                                                  Cmaj7 - Dmin7 - Fmin7 - D#maj7
EDM               i-bVI-bVII-V (Techno)                                                      Cmin - F - F# - E
EDM               I-V-vi-iii (Tropical House)                                                C - E - Fmin - Dmin
EDM               I-bVII-IV-I (Big Room House)                                               C - F# - D# - C
EDM               vi-V-IV-V (Electro House)                                                  Fmin - E - D# - E
EDM               I-iii-IV-vi (Chillstep)                                                    Cadd9 - Dmin7 - D#maj7 - Fmin7
EDM               i-bIII-bVII-IV (Drum and Bass)                                             Cmin - D - F# - D#
EDM               I-vi-IV-V-I-vi-IV-V (Extended EDM)                                         C - Fmin - D# - E - C - Fmin - D# - Esus4
Flamenco          i-VII-VI-V (Andalusian Cadence - Soleá, Bulerías)                          Cmin - F# - F - E
Flamenco          i-VII-iv-V (Variant Andalusian - Seguiriya)                                Cmin - F# - D#min - E
Flamenco          i-iv-V (Tango Flamenco)                                                    Cmin - D#min - E
Flamenco          i-VII-i-VII (Fandangos de Huelva)                                          Cmin - F# - Cmin - F#
Flamenco          i-v-VI-i (Phrygian Cadence - Soleá por Bulerías)                           Cmin - Emin - F - Cmin
Flamenco          I-V-I-V (Sevillanas)                                                       Cmaj - E - Cmaj - E
Flamenco          i-v-i-v (Tarantas)                                                         Cmin - Emin - Cmin - Emin
Flamenco          i-III-i-III (Farruca)                                                      Cmin - D - Cmin - D
Flamenco          i-VI-V-i (Bulerías al Golpe)                                               Cmin - F - E - Cmin
Flamenco          i7-iv7-V7-i7 (Jazz Flamenco)                                               Cmin7 - D#min7 - E7 - Cmin7
Flamenco          i-iv-V7-i (Rumba Flamenca)                                                 Cmin - D#min - E7 - Cmin
Flamenco          i-VII-VI-V-i (Extended Andalusian - Alegrías)                              Cmin - F# - F - E - Cmin
Flamenco          i-III-VI-i (Guajiras)                                                      Cmin - D - F - Cmin
Flamenco          i-iv-VII-VI (Peteneras)                                                    Cmin - D#min - F# - F
Flamenco          i-VI-VII-i (Tanguillos)                                                    Cmin - F - F# - Cmin
Flamenco          i-VII-VI-V-iv-III-II-I (Extended Phrygian - Granaína)                      Cmin - F# - F - E - D#min - D - C# - C
Flamenco          i-bII-bIII-i (Taranta Variant)                                             Cmin - C# - D - Cmin
Flamenco          i-VII-VI-v (Malagueña)                                                     Cmin - F# - F - Emin
Folk              I-IV-V (Basic Folk)                                                        C - D# - E
Folk              I-V-IV-I (Folk Rock)                                                       C - E - D# - C
Folk              I-vi-IV-V (50s Progression)                                                C - Fmin - D# - E
Folk              vi-IV-I-V (Modern Folk)                                                    Fmin - D# - C - E
Folk              I-V-vi-IV (Pop Folk)                                                       C - E - Fmin - D#
Folk              I-III-IV-I (Irish Folk)                                                    C - D - D# - C
Folk              i-VII-VI-V (Andalusian Cadence)                                            Cmin - F# - F - E
Folk              I-ii-iii-IV (Ascending Folk)                                               C - C#min - Dmin - D#
Folk              I-V-I-IV (Country Folk)                                                    C - E - C - D#
Folk              i-III-VII-iv (Minor Folk)                                                  Cmin - D - F# - D#min
Folk              I-IV-I-V (Bluegrass Turn)                                                  C - D# - C - E
Folk              I-bVII-IV-I (Modal Folk)                                                   C - F# - D# - C
Funk              I7-IV7 (Basic Funk)                                                        C7 - D#7
Funk              I9-IV9-V9 (Extended Funk)                                                  C9 - D#9 - E9
Funk              Im7-IV7-bVII7 (Minor Funk)                                                 Cmin7 - D#7 - F#7
Funk              I7-IV7-V7-IV7 (Funk Turnaround)                                            C7 - D#7 - E7 - D#7
Funk              I13-IV13-ii7-V7 (Jazz Funk)                                                C13 - D#13 - C#min7 - E7
Funk              Im7-bIII7-bVI7-V7 (Chromatic Funk)                                         Cmin7 - D7 - F7 - E7
Funk              I7-bVII7-IV7 (James Brown - I Got You (I Feel Good))                       C7 - F#7 - D#7
Funk              Im7-IV7 (Stevie Wonder - Superstition)                                     Cmin7 - D#7
Funk              I7-V7-IV7-I7 (Parliament - Give Up the Funk)                               C7 - E7 - D#7 - C7
Funk              Im7-bIII7-IV7-bII7 (Herbie Hancock - Chameleon)                            Cmin7 - D7 - D#7 - C#7
Funk              I9-IV9-V9-IV9 (Earth, Wind & Fire - September)                             C9 - D#9 - E9 - D#9
Funk              Im7-Vm7-IVm7 (Chic - Good Times)                                           Cmin7 - Emin7 - D#min7
Funk              I7-IV7-ii7-V7 (Tower of Power - What Is Hip?)                              C7 - D#7 - C#min7 - E7
Funk              Im9-bVII9-bIII9-bVI9 (Jamiroquai-style Funk)                               Cmin9 - F#9 - D9 - F9
Gospel            I-IV-V-IV (Basic Gospel)                                                   Cmaj7 - D#maj7 - E7 - D#maj7
Gospel            I-vi-ii-V (Gospel Turnaround)                                              Cmaj9 - Fmin7 - C#min7 - E9
Gospel            I-III-IV-V (Traditional Gospel)                                            Cmaj7 - D7 - D#maj7 - E7
Gospel            vi-IV-V-iii-ii-I (Contemporary Gospel)                                     Fmin7 - D#maj7 - E7 - Dmin7 - C#min7 - Cmaj9
Gospel            I-V-vi-IV (Modern Gospel)                                                  Cmaj9 - E9 - Fmin11 - D#maj9
Gospel            I-bIII-IV-iv (Gospel Blues)                                                C7 - D7 - D#maj7 - D#min7
Gospel            I-VI-ii-V (Jazz Gospel)                                                    Cmaj9 - F7 - C#min9 - E13
Gospel            I-V-vi-iii-IV-I-IV-V (Extended Gospel)                                     Cmaj7 - E7 - Fmin7 - Dmin7 - D#maj7 - Cmaj7 - D#maj7 - E7
Gospel            I-iii-IV-V (Soulful Gospel)                                                Cmaj9 - Dmin11 - D#maj9 - E13
Gospel            ii-V-I-vi (Gospel Jazz)                                                    C#min9 - E13 - Cmaj9 - Fmin11
Gospel            I-IV-V-III-vi-ii-V-I (Circle Progression)                                  Cmaj7 - D#maj7 - E7 - D7 - Fmin7 - C#min7 - E7 - Cmaj7
Gospel            I-V-vi-IV-I-V-iii-vi (Contemporary Worship)                                Cmaj7 - E7 - Fmin7 - D#maj7 - Cmaj7 - E7 - Dmin7 - Fmin7
Jazz              ii-V-I (Basic)                                                             C#min7 - E7 - Cmaj7
Jazz              ii-V-I (Extended)                                                          C#min7 - E7add13 - Cmaj7add9
Jazz              I-vi-ii-V (Basic)                                                          Cmaj7 - Fmin7 - C#min7 - E7
Jazz              Rhythm Changes (A Section)                                                 Cmaj7 - Cmaj7 - Fmin7 - E7 - C#min7 - E7 - Cmaj7 - Cmaj7
Jazz              Autumn Leaves (First 8 bars)                                               C#min7 - E7 - Cmaj7 - D#maj7 - F#min7b5 - A7 - Dmin - Dmin
Jazz              All The Things You Are (First 8 bars)                                      D#min7 - G#min7 - C#7 - F#maj7 - Bmin7 - Emin7 - A7 - Dmaj7
Jazz              Giant Steps (First 8 bars)                                                 Amaj7 - B7 - Emaj7 - F#7 - Bmaj7 - C#7 - F#maj7 - G#7
Jazz              Blue Bossa                                                                 Cmin - Cmin - D#min7 - G#7 - C#maj7 - C#maj7 - Fmin7b5 - A#7
Jazz              Coltrane Changes                                                           Cmaj7 - E7 - Amaj7 - C#7 - F#maj7 - A#7
Jazz              So What (Modal Jazz)                                                       C#minadd11 - C#minadd11 - C#minadd11 - C#minadd11 - Dminadd11 - Dminadd11 - C#minadd11 - C#minadd11
Jazz              Stella By Starlight (First 8 bars)                                         Dmin7 - G7 - Cmaj7 - Cmaj7 - D#min7 - G#7 - C#maj7 - C#maj7
Metal             I-VI-bVII-VI (Power Chord Progression)                                     C - F - F# - F
Metal             I-bVII-bVI-V (Chromatic Descent)                                           C - B - A# - A
Metal             i-bII-VI-iv (Phrygian Metal)                                               Cmin - C# - F - D#min
Metal             i-V-VI-i (Harmonic Minor Metal)                                            Cmin - E - F - Cmin
Metal             i-iii-IV-vi (Dorian Metal)                                                 Cmin - Dmin - D# - Fmin
Metal             I-bVI-I-bVI (Tritone-based Progression)                                    C - F# - C - F#
Metal             I-VI-VIII-VI (Black Sabbath - Iron Man)                                    C - F - G - F
Metal             i-IV-V-i (Metallica - Enter Sandman)                                       Cmin - D# - E - Cmin
Metal             i-bII-bVII-i (Slayer - Raining Blood)                                      Cmin - C# - B - Cmin
Metal             I-VI-bVII-I (Judas Priest - Breaking the Law)                              C - F - F# - C
Metal             i-vi-bII-I (Megadeth - Symphony of Destruction)                            Cmin - Fmin - G# - G
Metal             i-bVII-VI-IV (Dream Theater - Pull Me Under)                               Cmin - F# - F - D#
Metal             i-VI-IV-V (System of a Down - Chop Suey!)                                  Cmin - F - D# - E
Metal             i-bII-bVII-VI (Opeth - The Drapery Falls)                                  Cmin - G# - F# - F
Pop               I-V-vi-IV                                                                  C - E - Fmin - D#
Pop               vi-IV-I-V                                                                  Fmin - D# - C - E
Pop               I-vi-IV-V                                                                  C - Fmin - D# - E
Pop               I-IV-V-IV                                                                  C - D# - E - D#
Pop               ii-V-I-vi                                                                  C#min - E - C - Fmin
Pop               I-V-vi-iii-IV-I-IV-V                                                       C - E - Fmin - Dmin - D# - C - D# - E
Pop               I-V-vi-IV (Four Chords - Axis of Awesome)                                  C - E - Fmin - D#
Pop               vi-IV-I-V (Don't Stop Believin' - Journey)                                 Fmin - D# - C - E
Pop               I-V-vi-III-IV (Every Breath You Take - The Police)                         C - E - Fmin - D - D#
Pop               I-vi-IV-V (Stand By Me - Ben E. King)                                      C - Fmin - D# - E
Pop               I-V-vi-iii-IV-I-V (Despacito - Luis Fonsi & Daddy Yankee)                  C - E - Fmin - Dmin - D# - C - E
Pop               I-vi-IV-V (Achy Breaky Heart - Billy Ray Cyrus)                            C - Fmin - D# - E
Pop               vi-V-IV-V (Zombie - The Cranberries)                                       Fmin - E - D# - E
Pop               I-iii-vi-IV (Someone Like You - Adele)                                     C - Dmin - Fmin - D#
Pop               I-V-vi-IV-I-V-iii-IV (Wherever You Will Go - The Calling)                  C - E - Fmin - D# - C - E - Dmin - D#
Pop               I-V-vi-iii-IV-I-IV-V (Wonderwall - Oasis)                                  C - E - Fmin - Dmin - D# - C - D# - E
Progressive Rock  I-V-vi-iii-IV-I-IV-V                                                       C - E - Fmin - Dmin - D# - C - D# - E
Progressive Rock  i-bVII-bVI-bVII-i-bIII-bII-bIII                                            Cmin - F# - F - F# - Cmin - D - C# - D
Progressive Rock  I-IV-V-vi-ii-iii-IV-V                                                      Cmaj - D#maj - Emaj - Fmin - C#min - Dmin - D#maj - Emaj
Progressive Rock  I-bIII-IV-iv-bVI-V-I                                                       Cmaj - D - D#maj - D#min - F - Emaj - Cmaj
Progressive Rock  i-bVII-bVI-bVII (Starless - King Crimson)                                  Cmin - F# - F - F#
Progressive Rock  I-V-ii-IV (Roundabout - Yes)                                               Cmaj - Emaj - C#min - D#maj
Progressive Rock  i-bIII-bVII-IV-bVI (21st Century Schizoid Man - King Crimson)              Cmin - D - F# - D# - F
Progressive Rock  I-bVII-IV-I (Carry on Wayward Son - Kansas)                                C - F# - D# - C
Progressive Rock  i-bVI-bIII-bVII (Kashmir - Led Zeppelin)                                   Cmin - F - D - F#
Progressive Rock  I-IV-V-IV-bVII-IV (Tom Sawyer - Rush)                                      C - D# - E - D# - F# - D#
Progressive Rock  I-V-vi-iii-IV-I-V-vi (Firth of Fifth - Genesis)                            Cmaj - Emaj - Fmin - Dmin - D#maj - Cmaj - Emaj - Fmin
Progressive Rock  i-bVII-bVI-bVII-i-bIII-bII-bIII (Shine On You Crazy Diamond - Pink Floyd)  Cmin - F# - F - F# - Cmin - D - C# - D
Progressive Rock  I-ii-iii-IV-V-vi-vii° (Apocalypse in 9/8 - Genesis)                        Cmaj - C#min - Dmin - D#maj - Emaj - Fmin - F#dim
Progressive Rock  I-bIII-IV-bVI (Money - Pink Floyd)                                         C7 - D - D# - F
Punk              I-IV-V (Basic Punk)                                                        C - D# - E
Punk              I-V-vi-IV (Pop Punk)                                                       C - E - Fmin - D#
Punk              I-bVII-IV (Punk Rock)                                                      C - F# - D#
Punk              I-IV-bVII-IV (Alternative Punk)                                            C - D# - F# - D#
Punk              I-III-IV (Punk Power Progression)                                          C - D - D#
Punk              vi-IV-I-V (Melodic Punk)                                                   Fmin - D# - C - E
Punk              I-V-IV-IV (Ramones - Blitzkrieg Bop)                                       C - E - D# - D#
Punk              I-vi-IV-V (Green Day - Basket Case)                                        C - Fmin - D# - E
Punk              I-V-vi-iii-IV (Blink-182 - All The Small Things)                           C - E - Fmin - Dmin - D#
Punk              I-III-vi-IV (The Clash - Should I Stay or Should I Go)                     C - D - Fmin - D#
Punk              I-bVII-IV-I (Sex Pistols - Anarchy in the U.K.)                            C - F# - D# - C
Punk              I-V-vi-IV-I-V (NOFX - Linoleum)                                            C - E - Fmin - D# - C - E
Punk              I-IV-V-IV (Bad Religion - 21st Century Digital Boy)                        C - D# - E - D#
Punk              i-bVII-bVI-V (Minor Punk Progression)                                      Cmin - F# - F - E
Reggae            I-V (Basic Reggae)                                                         C - E
Reggae            I-IV-V (Roots Reggae)                                                      C - D# - E
Reggae            i-bVII-VI (Minor Reggae)                                                   Cmin - F# - F
Reggae            I-vi-IV-V (Ska Influence)                                                  C - Fmin - D# - E
Reggae            i-iv-V (Minor Roots)                                                       Cmin - D#min - E
Reggae            I-V-vi-IV (Pop Reggae)                                                     C - E - Fmin - D#
Reggae            I-bVII-bVI-V (Reggae Rock)                                                 C - F# - F - E
Reggae            i-bIII-bVII-i (Dub Reggae)                                                 Cmin - D - F# - Cmin
Reggae            I-IV-I-V (Reggae Fusion)                                                   C - D# - C - E
Reggae            i-bVI-bVII (Roots Rocksteady)                                              Cmin - F - F#
Reggae            I-iii-IV-I (Lover's Rock)                                                  C - Dmin - D# - C
Reggae            I-V-vi-IV-I-IV (Extended Reggae)                                           C - E - Fmin - D# - C - D#
R&B               ii-V-I-vi (Classic R&B)                                                    C#min7 - E7 - Cmaj7 - Fmin7
R&B               I-vi-IV-V (Doo-Wop Progression)                                            Cmaj7 - Fmin7 - D#maj7 - E7
R&B               i-iv-VII-III (Neo-Soul)                                                    Cmin7 - D#min7 - F#7 - Dmaj7
R&B               I-V-vi-IV (Contemporary R&B)                                               Cmaj7 - E7 - Fmin7 - D#maj7
R&B               ii-V-I-IV (Jazz-Influenced R&B)                                            C#min7 - E7 - Cmaj7 - D#maj7
R&B               vi-IV-I-V (90s R&B)                                                        Fmin7 - D#maj7 - Cmaj7 - E7
R&B               I-iii-IV-V (Soul Progression)                                              Cmaj7 - Dmin7 - D#maj7 - E7
R&B               i-bVII-bVI-V (Minor R&B)                                                   Cmin7 - F#7 - Fmaj7 - E7
R&B               I-vi-ii-V (Smooth R&B)                                                     Cmaj9 - Fmin9 - C#min9 - E9
R&B               I-V-vi-iii-IV-I-IV-V (Extended R&B)                                        Cmaj7 - E7 - Fmin7 - Dmin7 - D#maj7 - Cmaj7 - D#maj7 - E7
R&B               i-IV-VII-III (Alternative R&B)                                             Cmin7 - D#maj7 - F#7 - Dmaj7
R&B               I-bIII-IV-iv (Blues-Inspired R&B)                                          C7 - D7 - D#7 - D#min7
Rock              I-IV-V                                                                     C - D# - E
Rock              I-V-IV                                                                     C - E - D#
Rock              I-bVII-IV                                                                  C - F# - D#
Rock              I-vi-IV-V                                                                  C - Fmin - D# - E
Rock              i-bVII-bVI-V                                                               Cmin - F# - F - E
Rock              I-V-vi-IV                                                                  C - E - Fmin - D#
Rock              I-III-IV-IV                                                                C - D - D# - D#
Rock              I-II-IV-I                                                                  C - C# - D# - C
Rock              I-bVII-IV (Sweet Home Alabama - Lynyrd Skynyrd)                            C - F# - D#
Rock              I-V-vi-IV (When I Come Around - Green Day)                                 C - E - Fmin - D#
Rock              i-bVI-III-bVII (Smoke on the Water - Deep Purple)                          Cmin - F - D - F#
Rock              I-IV-V-IV (Louie Louie - The Kingsmen)                                     C - D# - E - D#
Rock              I-bIII-IV-IV (The Hardest Button to Button - The White Stripes)            C - D - D# - D#
Rock              I-V-IV-V (Twist and Shout - The Beatles)                                   C - E - D# - E
Rock              I-III-vi-IV (Under the Bridge - Red Hot Chili Peppers)                     C - D - Fmin - D#
Rock              V-IV-I (Pinball Wizard - The Who)                                          E - D# - C
Rock              i-bVII-bVI-bVII (Ain't Talkin' 'Bout Love - Van Halen)                     Cmin - F# - F - F#
Rock              I-V-vi-iii-IV-I-IV-V (Stairway to Heaven - Led Zeppelin)                   C - E - Fmin - Dmin - D# - C - D# - E
Rock              vi-IV-I-V (Zombie - The Cranberries)                                       Fmin - D# - C - E
Rock              I-vi-iii-IV-I-V (Where is My Mind - Pixies)                                C - Fmin - Dmin - D# - C - E
Salsa             I-IV-V-IV (Basic Salsa)                                                    C7 - D#7 - E7 - D#7
Salsa             i-iv-V7 (Minor Salsa)                                                      Cmin - D#min - E7
Salsa             I-V-vi-IV (Pop-influenced Salsa)                                           Cmaj - Emaj - Fmin - D#maj
Salsa             ii7-V7-I7 (Jazz-influenced Salsa)                                          C#min7 - E7 - C7
Salsa             I7-IV7-V7 (Son Montuno)                                                    C7 - D#7 - E7
Salsa             i-bVII-bVI-V (Andalusian Cadence in Salsa)                                 Cmin - F#maj - Fmaj - Emaj
Salsa             I-vi-ii-V (Salsa Romantica)                                                Cmaj7 - Fmin7 - C#min7 - E7
Salsa             i-IV-V-i (Héctor Lavoe - El Cantante)                                      Cmin - D#maj - Emaj - Cmin
Salsa             I7-IV7-V7 (Celia Cruz - La Vida Es Un Carnaval)                            C7 - D#7 - E7
Salsa             i-bVII-bVI-V7 (Marc Anthony - Valió la Pena)                               Cmin - F#maj - Fmaj - E7
Salsa             I-vi-IV-V (Rubén Blades - Pedro Navaja)                                    Cmaj - Fmin - D#maj - Emaj
Salsa             i-iv-V7-i (Willie Colón - Idilio)                                          Cmin - D#min - E7 - Cmin
Samba             I-ii-V7 (Basic Samba)                                                      Cmaj7 - C#min7 - E7
Samba             i7-iv7-V7 (Minor Samba)                                                    Cmin7 - D#min7 - E7
Samba             I7-IV7-V7 (Partido Alto Samba)                                             C7 - D#7 - E7
Samba             I-VI7-ii7-V7 (Samba Turnaround)                                            Cmaj7 - F7 - C#min7 - E7
Samba             I-III7-VI7-II7-V7 (Extended Samba)                                         Cmaj7 - D7 - F7 - C#7 - E7
Samba             Imaj7-vi7-ii7-V7 (Jobim - Samba de Uma Nota Só)                            Cmaj7 - Fmin7 - C#min7 - E7
Samba             i7-iv7-VII7-III7-VI7-ii7b5-V7 (Jorge Ben - Mas Que Nada)                   Cmin7 - D#min7 - F#7 - D7 - F7 - C#min7b5 - E7
Samba             I6-vi7-ii7-V7 (Ary Barroso - Aquarela do Brasil)                           Cadd13 - Fmin7 - C#min7 - E7
Samba             i-bVII-bVI-V7 (Cartola - As Rosas Não Falam)                               Cmin - F#maj - Fmaj - E7
Samba             Imaj7-vi7-iimin7-V7-iii7-VI7-ii7-V7 (João Bosco - Linha de Passe)          Cmaj7 - Fmin7 - C#min7 - E7 - Dmin7 - F7 - C#min7 - E7
Samba             I7-IV7-ii7-V7 (Zeca Pagodinho - Deixa a Vida Me Levar)                     C7 - D#7 - C#min7 - E7
Samba             i7-IV7-V7 (Paulinho da Viola - Timoneiro)                                  Cmin7 - D#7 - E7
Soul              I-IV-V (Basic Soul)                                                        Cmaj7 - D#maj7 - E7
Soul              ii-V-I (Jazz-influenced Soul)                                              C#min7 - E7 - Cmaj7
Soul              I-vi-IV-V (Doo-wop Soul)                                                   Cmaj7 - Fmin7 - D#maj7 - E7
Soul              i-iv-V (Minor Soul)                                                        Cmin7 - D#min7 - E7
Soul              I-iii-IV-V (Gospel-influenced Soul)                                        Cmaj7 - Dmin7 - D#maj7 - E7
Soul              I-V-vi-IV (Modern Soul)                                                    Cmaj9 - E9 - Fmin9 - D#maj9
Soul              I-bVII-IV-V (Otis Redding - (Sittin' On) The Dock of the Bay)              C - F# - D# - E
Soul              I-V-vi-IV (Sam Cooke - A Change Is Gonna Come)                             Cmaj7 - E7 - Fmin7 - D#maj7
Soul              i-bIII-IV-iv (Aretha Franklin - I Never Loved a Man (The Way I Love You))  Cmin - D - D# - D#min
Soul              I-IV-V-IV (Al Green - Let's Stay Together)                                 Cmaj9 - D#maj9 - E9 - D#maj9
Soul              i-iv-VII-VI (Marvin Gaye - What's Going On)                                Cmin9 - D#min9 - F#9 - Fmaj9
Soul              I-vi-ii-V (Stevie Wonder - Isn't She Lovely)                               Cmaj7 - Fmin7 - C#min7 - E7
Soul              i-IV-V-iv (Bill Withers - Ain't No Sunshine)                               Cmin - D# - E - D#min
Soul              I-V-vi-iii-IV-I-IV-V (D'Angelo - Untitled (How Does It Feel))              Cmaj9 - E9 - Fmin9 - Dmin9 - D#maj9 - Cmaj9 - D#maj9 - E9
```
</details>

---

Remember, `musegen` is a tool for exploration and inspiration. The true magic happens when you take these generated progressions and breathe life into them with your unique musical voice. Happy exploring!