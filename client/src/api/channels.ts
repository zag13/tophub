import http from "@/utils/http";

const prefix = "/api";
const channelsApi = {
  feed: `${prefix}/feed`,
};

export type FeedResult = {
  list: Feed[];
};

export type Feed = {
  site: string;
  rank: number;
  title: string;
  url: string;
  extra: {};
};

export const feed = (data?: { site?: string }) => {
  return http.request<FeedResult>({
    method: "post",
    url: channelsApi.feed,
    data: { site: data?.site },
  });
};
