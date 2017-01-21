<?php

namespace Tests\SubtitleAnalyzer;


use App\Src\SubtitleAnalyzer\SubtitleAnalyzer;

class SubtitleAnalyzerTest extends \PHPUnit_Framework_TestCase
{

    /**
     * @test
     * @dataProvider ProvideSubtitles
     *
     * @param $searchTerm
     * @param $expected
     * @param $subtitles
     */
    public function shouldReturnTimeforSubtitles($searchTerm, $expected, $subtitles)
    {
        $subtitleAnalyzer = new SubtitleAnalyzer();
        $expected = $subtitleAnalyzer->getTiming($subtitles, $searchTerm);
    }

    public function ProvideSubtitles()
    {
        return [
            [
                'searchKeyword' => '',
                'expected' => [
                    [
                        0 => [
                            "_index" => "videos_en",
                            "_type" => "videosSubtitles",
                            "_id" => "n4bucphC9r4",
                            "_score" => 1.756412,
                            "_source" => [
                                "start" => 12,
                                "sentence" => ",align:start position:19%\npasta<c.colorCCCCCC><00:00:02.879><c> hot
</c><00:00:03.330><c>dog</c><00:00:03.720><c> what</c><00:00:03.840><c> better</c><00:00:04.140><c> way</c><00:00:04.290>
<c> to</c><00:00:04.350><c> end</c></c>",
                                "subtitles" => [
                                    "image_default" => "https://i.ytimg.com/vi/n4bucphC9r4/default.jpg",
                                    "image_medium" => "https://i.ytimg.com/vi/n4bucphC9r4/mqdefault.jpg",
                                    "_subtitles" => [
                                        "sentence" => [
                                            0 => "align:start position:19%\n<c.colorE5E5E5>hear<00:00:00.240><c> me</c>
<00:00:00.390><c> out</c><00:00:00.930><c> we've</c><00:00:01.140><c> had</c><00:00:01.319><c> steak</c><00:00:02.129>
<c> we've</c><00:00:02.250><c> had</c></c>",
                                            1 => ",align:start position:19%\npasta<c.colorCCCCCC><00:00:02.879><c> hot
</c><00:00:03.330><c>dog</c><00:00:03.720><c> what</c><00:00:03.840><c> better</c><00:00:04.140><c> way</c><00:00:04.290>
<c> to</c><00:00:04.350><c> end</c></c>",
                                        ],
                                        "start" => [
                                            0 => "00:00:00.000",
                                            1 => "00:00:02.370",
                                            2 => "00:00:04.500",
                                        ],
                                        "end" => [],
                                    ],
                                    "video_hash_id" => "n4bucphC9r4",
                                    "language" => "en",
                                    "tags" => null,
                                    "video_url" => "https://www.youtube.com/watch?v=n4bucphC9r4",
                                    "@timestamp" => "2017-01-21T09:06:12.516Z",
                                    "video_title" => "$27 Cake Vs. $1,120 Cake",
                                    "@version" => "1",
                                    "id" => 24,
                                    "image_high" => "https://i.ytimg.com/vi/n4bucphC9r4/hqdefault.jpg",
                                    "views" => null,
                                    "video_id" => 68,
                                    "upload_date" => null,
                                ],
                            ],
                        ],
                    ],
                ],
                'subtitles' => [
                    [
                        0 => [
                            "_index" => "videos_en",
                            "_type" => "videosSubtitles",
                            "_id" => "n4bucphC9r4",
                            "_score" => 1.756412,
                            "_source" => [
                                "subtitles" => [
                                    "image_default" => "https://i.ytimg.com/vi/n4bucphC9r4/default.jpg",
                                    "image_medium" => "https://i.ytimg.com/vi/n4bucphC9r4/mqdefault.jpg",
                                    "_subtitles" => [
                                        "sentence" => [
                                            0 => "align:start position:19%\n<c.colorE5E5E5>hear<00:00:00.240><c> me</c><00:00:00.390><c> out</c><00:00:00.930><c> we've</c><00:00:01.140><c> had</c><00:00:01.319><c> steak</c><00:00:02.129><c> we've</c><00:00:02.250><c> had</c></c>",
                                            1 => ",align:start position:19%\npasta<c.colorCCCCCC><00:00:02.879><c> hot</c><00:00:03.330><c>dog</c><00:00:03.720><c> what</c><00:00:03.840><c> better</c><00:00:04.140><c> way</c><00:00:04.290><c> to</c><00:00:04.350><c> end</c></c>",
                                        ],
                                        "start" => [
                                            0 => "00:00:00.000",
                                            1 => "00:00:02.370",
                                            2 => "00:00:04.500",
                                        ],
                                        "end" => [],
                                    ],
                                    "video_hash_id" => "n4bucphC9r4",
                                    "language" => "en",
                                    "tags" => null,
                                    "video_url" => "https://www.youtube.com/watch?v=n4bucphC9r4",
                                    "@timestamp" => "2017-01-21T09:06:12.516Z",
                                    "video_title" => "$27 Cake Vs. $1,120 Cake",
                                    "@version" => "1",
                                    "id" => 24,
                                    "image_high" => "https://i.ytimg.com/vi/n4bucphC9r4/hqdefault.jpg",
                                    "views" => null,
                                    "video_id" => 68,
                                    "upload_date" => null,
                                ],
                            ],
                        ],
                        1 => [
                            "_index" => "videos_en",
                            "_type" => "videosSubtitles",
                            "_id" => "n4bucphC9r4",
                            "_score" => 1.756412,
                            "_source" => [
                                "subtitles" => [
                                    "image_default" => "https://i.ytimg.com/vi/n4bucphC9r4/default.jpg",
                                    "image_medium" => "https://i.ytimg.com/vi/n4bucphC9r4/mqdefault.jpg",
                                    "_subtitles" => [
                                        "sentence" => [
                                            0 => "bluh bluh bluh",
                                            1 => "almost a better way to ",
                                        ],
                                        "start" => [
                                            0 => "00:00:00.000",
                                            1 => "00:00:02.370",
                                            2 => "00:00:04.500",
                                        ],
                                        "end" => [],
                                    ],
                                    "video_hash_id" => "n4bucphC9r4",
                                    "language" => "en",
                                    "tags" => null,
                                    "video_url" => "https://www.youtube.com/watch?v=n4bucphC9r4",
                                    "@timestamp" => "2017-01-21T09:06:12.516Z",
                                    "video_title" => "$27 Cake Vs. $1,120 Cake",
                                    "@version" => "1",
                                    "id" => 24,
                                    "image_high" => "https://i.ytimg.com/vi/n4bucphC9r4/hqdefault.jpg",
                                    "views" => null,
                                    "video_id" => 68,
                                    "upload_date" => null,
                                ],
                            ],
                        ],
                    ],
                ],
            ]
    }
}