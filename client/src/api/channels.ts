import http from "@/utils/http";

export type FeedResult = {
  list: [Feed];
};

export type Feed = {
  site: string;
  ranking: number;
  title: string;
  url: string;
};

export const feed = (data?: { site?: string }) => {
  return http.request<FeedResult>({
    method: "post",
    url: "/api/feed",
    data: { site: data?.site },
  });
};
